package subcommands

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	gent "github.com/Creaft-JP/tit/db/global/ent"
	"github.com/Creaft-JP/tit/db/global/ent/globalconfig"
	"github.com/Creaft-JP/tit/db/global/ent/logintoken"
	lent "github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/page"
	"github.com/Creaft-JP/tit/db/local/ent/remote"
	"github.com/Creaft-JP/tit/db/local/ent/section"
	"github.com/Creaft-JP/tit/db/local/ent/titcommit"
	e "github.com/Creaft-JP/tit/error"
	reqi "github.com/Creaft-JP/tit/json/request/images"
	"github.com/Creaft-JP/tit/json/request/mb"
	"github.com/Creaft-JP/tit/json/response"
	resi "github.com/Creaft-JP/tit/json/response/images"
	"github.com/google/uuid"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func Push(args []string, gcl *gent.Client, lcl *lent.Client, ctx context.Context) (ret error) {
	if len(args) != 1 {
		return failure.New(e.Operation, failure.Message("subcommand push requires 1 argument"))
	}
	rem, err := lcl.Remote.Query().Where(remote.Name(args[0])).Only(ctx)
	if err != nil {
		if lent.IsNotFound(err) {
			return failure.Translate(err, e.Operation, failure.Messagef("remote \"%s\" doesn't exist", args[0]))
		} else {
			return failure.Translate(err, e.Database)
		}
	}
	av, err := generateAuthorizationValue(gcl, ctx)
	if err != nil {
		return failure.Wrap(err)
	}
	if err := uploadMainBody(rem.URL, av, lcl, ctx); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
func generateAuthorizationValue(gcl *gent.Client, ctx context.Context) (string, error) {
	user, err := gcl.GlobalConfig.Query().
		Where(globalconfig.Key("default-sign-in-user-slug")).Only(ctx)
	if err != nil {
		if gent.IsNotFound(err) {
			return "", failure.Translate(
				err,
				e.Operation,
				failure.Message("please login as a user who has permission to write to the remote repository"),
			)
		} else {
			return "", failure.Translate(err, e.Database)
		}
	}
	token, err := gcl.LoginToken.Query().Where(logintoken.SignInUserSlug(user.Value)).Only(ctx)
	if err != nil {
		if gent.IsNotFound(err) {
			return "", failure.Translate(
				err,
				e.Operation,
				failure.Message("please login as a user who has permission to write to the remote repository"),
			)
		} else {
			return "", failure.Translate(err, e.Database)
		}
	}
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", user.Value, token.CliLoginToken)),
	)), nil
}
func uploadImages(rem string, av string, lcl *lent.Client, ctx context.Context) (ret error) {
	root := reqi.Root{Resources: []reqi.Resource{}}
	resources, err := lcl.Image.Query().All(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	for _, resource := range resources {
		root.Resources = append(root.Resources, reqi.Resource{Id: resource.ID, Extension: resource.Extension, Description: resource.Description})
	}
	body, _ := json.Marshal(root)
	lreq, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/images", rem),
		bytes.NewReader(body),
	)
	lreq.Header.Set("Authorization", av)
	lreq.Header.Set("Content-Type", "application/json")
	lres, err := http.DefaultClient.Do(lreq)
	if err != nil {
		return failure.Translate(err, e.Network)
	}
	defer func(rc io.ReadCloser) {
		ret = multierr.Append(ret, failure.Translate(rc.Close(), e.Network))
	}(lres.Body)
	if err := getErrorFrom(lres); err != nil {
		return failure.Wrap(err)
	}
	decoder := json.NewDecoder(lres.Body)
	var lacks resi.Root
	if err := decoder.Decode(&lacks); err != nil {
		return failure.Translate(err, e.Network)
	}
	for _, i := range lacks.RequiredResourceImages {
		model, err := lcl.Image.Get(ctx, i.Id)
		if err != nil {
			return failure.Translate(err, e.Database)
		}
		ureq, err := http.NewRequest("POST", i.UploadUrl, generateMultipartBody(model))
		ures, err := http.DefaultClient.Do(ureq)
		if err != nil {
			return failure.Translate(err, e.Network)
		}
		defer func(rc io.ReadCloser) {
			ret = multierr.Append(ret, failure.Translate(rc.Close(), e.Network))
		}(ures.Body)
		if ures.StatusCode != 200 {
			return failure.New(e.Network, failure.Message(ures.Status))
		}
	}
	return nil
}
func generateMultipartBody(mo *lent.Image) io.Reader {
	buffer := bytes.NewBuffer([]byte{})
	writer := multipart.NewWriter(buffer)
	defer func(w *multipart.Writer) {
		_ = w.Close()
	}(writer)
	file, _ := writer.CreateFormFile("file", filepath.Base(mo.Description))
	_, _ = io.Copy(file, bytes.NewReader(mo.Contents))
	return buffer
}
func uploadMainBody(rem string, av string, lcl *lent.Client, ctx context.Context) (ret error) {
	root := mb.Root{Repository: mb.Repository{Pages: []mb.Page{}}}
	pages, err := lcl.Page.Query().Order(page.ByNumber()).All(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	for _, mpage := range pages {
		sections, err := mpage.QuerySections().Order(section.ByNumber()).All(ctx)
		if err != nil {
			return failure.Translate(err, e.Database)
		}
		jpage := mb.Page{Pathname: mpage.Pathname, Title: mpage.Title, ResourceIds: []uuid.UUID{}, Sections: []mb.Section{}}
		for _, msection := range sections {
			commits, err := msection.QueryCommits().Order(titcommit.ByNumber()).All(ctx)
			if err != nil {
				return failure.Translate(err, e.Database)
			}
			jsection := mb.Section{Slug: msection.Slug, Title: msection.Title, ResourceIds: []uuid.UUID{}, Commits: []mb.Commit{}}
			for _, mcommit := range commits {
				files, err := mcommit.QueryFiles().All(ctx)
				if err != nil {
					return failure.Translate(err, e.Database)
				}
				images, err := mcommit.QueryImages().All(ctx)
				if err != nil {
					return failure.Translate(err, e.Database)
				}
				jcommit := mb.Commit{Message: mcommit.Message, Files: []mb.File{}, ResourceIds: []uuid.UUID{}}
				for _, file := range files {
					jcommit.Files = append(jcommit.Files, mb.File{Path: file.Path, Content: file.Content})
				}
				for _, image := range images {
					jcommit.ResourceIds = append(jcommit.ResourceIds, image.ID)
				}
				jsection.Commits = append(jsection.Commits, jcommit)
			}
			jpage.Sections = append(jpage.Sections, jsection)
		}
		root.Repository.Pages = append(root.Repository.Pages, jpage)
	}
	body, err := json.Marshal(root)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", rem, bytes.NewBuffer(body))
	req.Header.Set("Authorization", av)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return failure.Translate(err, e.Network)
	}
	defer func(r *http.Response) {
		ret = multierr.Append(ret, r.Body.Close())
	}(res)
	return failure.Wrap(getErrorFrom(res))
}
func getErrorFrom(re *http.Response) error {
	if re.StatusCode == 204 {
		return nil
	} else {
		decoder := json.NewDecoder(re.Body)
		var body response.Error
		if err := decoder.Decode(&body); err != nil {
			return failure.Translate(err, e.Network)
		}
		return failure.New(
			e.Network,
			failure.Messagef(fmt.Sprintf("%s\nreason: %s", re.Status, body.Reason)),
		)
	}
}
