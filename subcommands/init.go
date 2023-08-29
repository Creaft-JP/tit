package subcommands

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/Creaft-JP/tit/ent"
	e "github.com/Creaft-JP/tit/error"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
)

func Init() (ret error) {
	if err := checkAlreadyInitialized(); err != nil {
		return failure.Wrap(err)
	}
	client, err := ent.Open(dialect.SQLite, "./.tit?_fk=1")
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	defer func(client *ent.Client) {
		ret = multierr.Append(err, failure.Translate(client.Close(), e.File))
	}(client)

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return failure.Translate(err, e.Database)
	}
	return nil
}

func checkAlreadyInitialized() error {
	_, err := os.Stat("./tit")
	if err == nil {
		return failure.New(e.Operation, failure.Message("tit repository already exists"))
	}
	if os.IsNotExist(err) {
		return nil
	}
	return failure.Translate(err, e.File)
}
