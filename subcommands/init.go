package subcommands

import (
	"context"
	"database/sql"
	"github.com/Creaft-JP/tit/db/local"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/directories"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/skeleton"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
)

func Init(ctx context.Context) (ret error) {
	isInitialized, err := directories.Exists(skeleton.Path)
	if err != nil {
		return failure.Wrap(err)
	}
	if isInitialized {
		return failure.New(e.Operation, failure.Message("tit repository already exists"))
	}
	if err := os.Mkdir(skeleton.Path, 0755); err != nil {
		return failure.Translate(err, e.File)
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
	transaction, err := client.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	section, err := transaction.Section.Create().
		SetSlug("section").
		SetNumber(1).
		SetTitle("").
		SetOverviewSentence("").Save(ctx)
	if err != nil {
		return multierr.Append(failure.Translate(err, e.Database), transaction.Rollback())
	}
	if _, err := transaction.Page.Create().
		SetPathname("/").
		SetNumber(1).
		SetTitle("").
		SetOverviewSentence("").
		AddSections(section).Save(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), transaction.Rollback())
	}
	return failure.Translate(transaction.Commit(), e.Database)
}
