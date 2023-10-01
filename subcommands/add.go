package subcommands

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
	"path/filepath"
)

func Add(args []string, cl *ent.Client, ctx context.Context) (ret error) {
	var files []file
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			if os.IsNotExist(err) {
				return failure.Translate(err, e.Operation, failure.Messagef("path %s doesn't exist", arg))
			} else {
				return failure.Translate(err, e.File)
			}
		}
		if info.IsDir() {
			return failure.New(e.Operation, failure.Messagef("path %s is a directory", arg))
		}
		content, err := os.ReadFile(arg)
		if err != nil {
			return failure.Translate(err, e.File)
		}
		wd, err := os.Getwd()
		if err != nil {
			return failure.Translate(err, e.File)
		}
		path, err := filepath.Abs(arg)
		if err != nil {
			return failure.Translate(err, e.File)
		}
		path, err = filepath.Rel(wd, path)
		if err != nil {
			return failure.Translate(err, e.File)
		}
		if path[:3] == fmt.Sprintf("..%c", filepath.Separator) {
			return failure.New(e.Operation, failure.Message("can't add a file out of tit repository"))
		}
		path = filepath.ToSlash(path)
		files = append(files, file{path: path, content: string(content)})
	}
	transaction, err := cl.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	if _, err := transaction.StagedFile.Delete().Where(stagedfile.PathIn(args...)).Exec(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), failure.Translate(transaction.Rollback(), e.Database))
	}
	var txFiles []*ent.StagedFileCreate
	for _, f := range files {
		txFiles = append(txFiles, transaction.StagedFile.Create().SetPath(f.path).SetContent(f.content))
	}
	if _, err := transaction.StagedFile.CreateBulk(txFiles...).Save(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), failure.Translate(transaction.Rollback(), e.Database))
	}
	return failure.Translate(transaction.Commit(), e.Database)
}

type file struct {
	path    string
	content string
}
