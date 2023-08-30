package main

import (
	"context"
	"github.com/Creaft-JP/tit/db"
	"github.com/Creaft-JP/tit/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/subcommands"
	"github.com/Creaft-JP/tit/subcommands/remote"
	"github.com/morikuni/failure"
	"os"
)

func main() {
	args := os.Args
	ctx := context.Background()

	client, err := db.MakeClient(db.FilePath)
	if err != nil {
		e.Handle(err)
		return
	}
	defer func(client *ent.Client) {
		e.Handle(failure.Translate(client.Close(), e.Database))
	}(client)

	if err := db.Migrate(client, ctx); err != nil {
		e.Handle(err)
		return
	}

	if err := route(args[1:], client, ctx); err != nil {
		e.Handle(err)
		return
	}
}

func route(args []string, client *ent.Client, ctx context.Context) error {
	if len(args) == 0 {
		return failure.New(e.Operation, failure.Message("subcommand must be specified"))
	}
	switch args[0] {
	case "init":
		return failure.Wrap(initRoute(ctx))
	case "remote":
		return failure.Wrap(remoteRoute(args[1:], client, ctx))
	default:
		return failure.New(e.Operation, failure.Messagef("subcommand: \"%s\" does not exits", args[0]))
	}
}

func initRoute(ctx context.Context) (err error) {
	return failure.Wrap(subcommands.Init(ctx))
}

func remoteRoute(args []string, client *ent.Client, ctx context.Context) (err error) {
	if len(args) > 0 {
		switch args[0] {
		case "add":
			return failure.Wrap(remoteAddRoute(args[1:], client, ctx))
		}
	}
	return failure.Wrap(subcommands.Remote(args, os.Stdout, client, ctx))
}

func remoteAddRoute(args []string, client *ent.Client, ctx context.Context) (err error) {
	return failure.Wrap(remote.Add(args, client, ctx))
}
