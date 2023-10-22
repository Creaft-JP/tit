package subcommands

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"flag"
	"fmt"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/page"
	"github.com/Creaft-JP/tit/db/local/ent/section"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/skeleton"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const commitMessageDefault = "_tit_commit_message_default_value"

func Commit(args []string, cl *ent.Client, ctx context.Context) error {
	fs := flag.NewFlagSet("commit", flag.ContinueOnError)
	message := fs.String("m", commitMessageDefault, "commit message")
	if err := parse(fs, args); err != nil {
		return failure.Wrap(err)
	}
	if *message == commitMessageDefault {
		editor, ok := os.LookupEnv("EDITOR")
		if !ok {
			editor = "vi"
		}
		tmp := filepath.Join(skeleton.Path, fmt.Sprintf("commit message (%d)", os.Getpid()))
		file, err := os.Create(tmp)
		if err != nil {
			return failure.Translate(err, e.File)
		}
		if err := file.Close(); err != nil {
			return failure.Translate(err, e.File)
		}
		command := exec.Command(editor, tmp)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			return failure.Translate(err, e.Editor)
		}
		content, err := os.ReadFile(tmp)
		if err != nil {
			return failure.Translate(err, e.File)
		}
		*message = strings.TrimSpace(string(content))
		if err := os.Remove(tmp); err != nil {
			return failure.Translate(err, e.File)
		}
	}
	return failure.Wrap(commit(*message, cl, ctx))
}
func commit(me string, cl *ent.Client, ctx context.Context) error {
	if me == "" {
		return failure.New(e.Operation, failure.Message("aborting commit due to empty commit message"))
	}
	pag, err := cl.Page.Query().Order(page.ByNumber(sql.OrderDesc())).First(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	sect, err := pag.QuerySections().Order(section.ByNumber(sql.OrderDesc())).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return failure.Translate(err, e.Operation, failure.Messagef("please start first section of page %s", pag.Pathname))
		} else {
			return failure.Translate(err, e.Database)
		}
	}
	number, err := sect.QueryCommits().Count(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	number += 1
	sf, err := cl.StagedFile.Query().All(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	transaction, err := cl.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	var builders []*ent.CommittedFileCreate
	for _, file := range sf {
		builders = append(builders, transaction.CommittedFile.Create().SetPath(file.Path).SetContent(file.Content))
	}
	cf, err := transaction.CommittedFile.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return multierr.Append(failure.Translate(err, e.Database), transaction.Rollback())
	}
	commit, err := transaction.TitCommit.Create().SetNumber(number).SetMessage(me).AddFiles(cf...).Save(ctx)
	if err != nil {
		return multierr.Append(failure.Translate(err, e.Database), transaction.Rollback())
	}
	if _, err := transaction.Section.UpdateOne(sect).AddCommits(commit).Save(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), transaction.Rollback())
	}
	if _, err := transaction.StagedFile.Delete().Exec(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), transaction.Rollback())
	}
	return transaction.Commit()
}
