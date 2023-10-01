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
	presentation "github.com/Creaft-JP/tit/json"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"net/http"
)

func Push(args []string, gcl *gent.Client, lcl *lent.Client, ctx context.Context) (ret error) {
	if len(args) != 1 {
		return failure.New(e.Operation, failure.Message("subcommand push requires 1 argument"))
	}
	root := presentation.Root{Repository: presentation.Repository{Pages: []presentation.Page{}}}
	pages, err := lcl.Page.Query().Order(page.ByNumber()).All(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	for _, mpage := range pages {
		sections, err := mpage.QuerySections().Order(section.ByNumber()).All(ctx)
		if err != nil {
			return failure.Translate(err, e.Database)
		}
		jpage := presentation.Page{Pathname: mpage.Pathname, Title: mpage.Title, Sections: []presentation.Section{}}
		for _, msection := range sections {
			commits, err := msection.QueryCommits().Order(titcommit.ByNumber()).All(ctx)
			if err != nil {
				return failure.Translate(err, e.Database)
			}
			jsection := presentation.Section{Slug: msection.Slug, Title: msection.Title, Commits: []presentation.Commit{}}
			for _, mcommit := range commits {
				files, err := mcommit.QueryFiles().All(ctx)
				if err != nil {
					return failure.Translate(err, e.Database)
				}
				jcommit := presentation.Commit{Message: mcommit.Message, Files: []presentation.File{}}
				for _, file := range files {
					jcommit.Files = append(jcommit.Files, presentation.File{Path: file.Path, Content: file.Content})
				}
				jsection.Commits = append(jsection.Commits, jcommit)
			}
			jpage.Sections = append(jpage.Sections, jsection)
		}
		root.Repository.Pages = append(root.Repository.Pages, jpage)
	}
	rem, err := lcl.Remote.Query().Where(remote.Name(args[0])).Only(ctx)
	if err != nil {
		if lent.IsNotFound(err) {
			return failure.Translate(err, e.Operation, failure.Messagef("remote \"%s\" doesn't exist", args[0]))
		} else {
			return failure.Translate(err, e.Database)
		}
	}
	body, err := json.Marshal(root)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("PUT", rem.URL, bytes.NewBuffer(body))
	user, err := gcl.GlobalConfig.Query().
		Where(globalconfig.Key("default-sign-in-user-slug")).Only(ctx)
	if err != nil {
		if gent.IsNotFound(err) {
			return failure.Translate(
				err,
				e.Operation,
				failure.Message("please login as a user who has permission to write to the remote repository"),
			)
		} else {
			return failure.Translate(err, e.Database)
		}
	}
	token, err := gcl.LoginToken.Query().Where(logintoken.SignInUserSlug(user.Value)).Only(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	request.Header.Set(
		"Authorization",
		fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s", user.Value, token.CliLoginToken)),
		)),
	)
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return failure.Translate(err, e.Network)
	}
	defer func(r *http.Response) {
		ret = multierr.Append(ret, r.Body.Close())
	}(response)
	if response.StatusCode != 204 {
		decoder := json.NewDecoder(response.Body)
		body := presentation.Error{Error: false, Reason: ""}
		if err := decoder.Decode(&body); err != nil {
			return failure.Translate(err, e.Network)
		}
		return failure.New(
			e.Network,
			failure.Messagef(fmt.Sprintf("%s\nreason: %s", response.Status, body.Reason)),
		)
	}
	return nil
}
