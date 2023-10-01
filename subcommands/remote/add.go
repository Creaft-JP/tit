package remote

import (
	"context"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/remote"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
)

func Add(args []string, client *ent.Client, ctx context.Context) error {

	if len(args) != 2 {
		return failure.New(e.Operation, failure.Messagef("args should be 2, but received %d", len(args)))
	}

	name := args[0]
	url := args[1]

	_, err := client.Remote.Query().Where(remote.Name(name)).Only(ctx)
	if !ent.IsNotFound(err) {
		if err != nil {
			return failure.Translate(err, e.Database)
		}
		return failure.New(e.Operation, failure.Messagef("remote %s already exists", name))
	}

	if _, err := client.Remote.Create().SetName(name).SetURL(url).Save(ctx); err != nil {
		return failure.Translate(err, e.Database)
	}

	return nil
}
