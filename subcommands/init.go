package subcommands

import (
	"context"
	"github.com/Creaft-JP/tit/db"
	"github.com/Creaft-JP/tit/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/skelton"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
)

func Init(ctx context.Context) (ret error) {
	isInitialized, err := skelton.IsAlreadyInitialized(skelton.Path)
	if err != nil {
		return failure.Wrap(err)
	}
	if isInitialized {
		return failure.New(e.Operation, failure.Message("tit repository already exists"))
	}
	if err := os.Mkdir(skelton.Path, 0755); err != nil {
		return failure.Translate(err, e.File)
	}
	client, err := db.MakeClient(db.FilePath)
	if err != nil {
		return failure.Wrap(err)
	}
	defer func(client *ent.Client) {
		ret = multierr.Append(err, failure.Translate(client.Close(), e.Database))
	}(client)

	if err := db.Migrate(client, ctx); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
