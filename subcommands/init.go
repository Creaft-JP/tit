package subcommands

import (
	"context"
	"github.com/Creaft-JP/tit/db/local"
	"github.com/Creaft-JP/tit/db/local/ent"
	e "github.com/Creaft-JP/tit/error"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
)

func Init(ctx context.Context) (ret error) {
	isInitialized, err := local.IsAlreadyInitialized(local.FilePath)
	if err != nil {
		return failure.Wrap(err)
	}
	if isInitialized {
		return failure.New(e.Operation, failure.Message("tit repository already exists"))
	}
	client, err := local.MakeClient(local.FilePath)
	if err != nil {
		return failure.Wrap(err)
	}
	defer func(client *ent.Client) {
		ret = multierr.Append(err, failure.Translate(client.Close(), e.Database))
	}(client)

	if err := local.Migrate(client, ctx); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
