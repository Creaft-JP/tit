package remote

import (
	"context"
	"github.com/Creaft-JP/tit/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
)

func Add(args []string, client *ent.Client, ctx context.Context) error {

	if len(args) != 2 {
		return failure.New(e.Operation, failure.Messagef("args should be 2, but received %d", len(args)))
	}

	name := args[0]
	url := args[1]

	if _, err := client.Remote.Create().SetName(name).SetURL(url).Save(ctx); err != nil {
		return failure.Translate(err, e.Database)
	}

	return nil
}
