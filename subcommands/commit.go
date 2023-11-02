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
	"github.com/mattn/go-shellwords"
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
			return failure.New(e.Operation, failure.Message("please specify EDITOR what you edit the message on"))
		}
		tmp := filepath.Join(skeleton.Path, fmt.Sprintf("commit message (%d)", os.Getpid()))
		line := fmt.Sprintf("%s '%s'", editor, tmp)
		environments, words, err := shellwords.ParseWithEnvs(line)
		if err != nil {
			return failure.Translate(err, e.Operation, failure.Messagef("failed to parse \"%s\" to shell words", line))
		}
		command := exec.Command(words[0], words[1:]...)
		command.Env = append(command.Environ(), environments...)
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
	return failure.Wrap(commit(*message, nil, cl, ctx))
}
func commit(me string, re []string, cl *ent.Client, ctx context.Context) error {
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
	var cfc []*ent.CommittedFileCreate
	for _, file := range sf {
		cfc = append(cfc, transaction.CommittedFile.Create().SetPath(file.Path).SetContent(file.Content))
	}
	cf, err := transaction.CommittedFile.CreateBulk(cfc...).Save(ctx)
	if err != nil {
		return multierr.Append(
			failure.Translate(err, e.Database),
			failure.Translate(transaction.Rollback(), e.Database),
		)
	}
	var ic []*ent.ImageCreate
	for i, in := range re {
		status, err := os.Stat(in)
		if err != nil {
			if os.IsNotExist(err) {
				return multierr.Append(
					failure.Translate(err, e.Operation, failure.Messagef("image \"%s\" isn't found", in)),
					failure.Translate(transaction.Rollback(), e.Database),
				)
			} else {
				return multierr.Append(
					failure.Translate(err, e.File),
					failure.Translate(transaction.Rollback(), e.Database),
				)
			}
		}
		if status.IsDir() {
			return multierr.Append(
				failure.New(e.Operation, failure.Messagef("image \"%s\" is a directory", in)),
				failure.Translate(transaction.Rollback(), e.Database),
			)
		}
		extension := filepath.Ext(in)
		if extension == "" {
			return multierr.Append(
				failure.New(e.Operation, failure.Messagef("image \"%s\" doesn't have an extension", in)),
				failure.Translate(transaction.Rollback(), e.Database),
			)
		}
		contents, err := os.ReadFile(in)
		if err != nil {
			return multierr.Append(
				failure.Translate(err, e.File),
				failure.Translate(transaction.Rollback(), e.Database),
			)
		}
		ic = append(ic, transaction.Image.Create().SetNumber(i+1).SetExtension(extension).SetContents(contents).SetDescription(in))
	}
	i, err := transaction.Image.CreateBulk(ic...).Save(ctx)
	if err != nil {
		return multierr.Append(
			failure.Translate(err, e.Database),
			failure.Translate(transaction.Rollback(), e.Database),
		)
	}
	commit, err := transaction.TitCommit.Create().SetNumber(number).SetMessage(me).AddFiles(cf...).AddImages(i...).Save(ctx)
	if err != nil {
		return multierr.Append(
			failure.Translate(err, e.Database),
			failure.Translate(transaction.Rollback(), e.Database),
		)
	}
	if _, err := transaction.Section.UpdateOne(sect).AddCommits(commit).Save(ctx); err != nil {
		return multierr.Append(
			failure.Translate(err, e.Database),
			failure.Translate(transaction.Rollback(), e.Database),
		)
	}
	if _, err := transaction.StagedFile.Delete().Exec(ctx); err != nil {
		return multierr.Append(
			failure.Translate(err, e.Database),
			failure.Translate(transaction.Rollback(), e.Database),
		)
	}
	return transaction.Commit()
}
