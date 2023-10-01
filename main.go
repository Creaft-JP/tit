package main

import (
	"context"
	gdb "github.com/Creaft-JP/tit/db/global"
	gent "github.com/Creaft-JP/tit/db/global/ent"
	"github.com/Creaft-JP/tit/db/local"
	lent "github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/directories"
	e "github.com/Creaft-JP/tit/error"
	g "github.com/Creaft-JP/tit/global"
	"github.com/Creaft-JP/tit/skeleton"
	"github.com/Creaft-JP/tit/subcommands"
	"github.com/Creaft-JP/tit/subcommands/remote"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
)

func main() {
	args := os.Args
	ctx := context.Background()

	initialized, err := directories.Exists(g.Path)
	if err != nil {
		e.Handle(err)
		return
	}
	if !initialized {
		if err := os.Mkdir(g.Path, 0755); err != nil {
			e.Handle(failure.Translate(err, e.File))
			return
		}
	}
	client, err := gdb.MakeClient(gdb.FilePath)
	defer func(c *gent.Client) {
		if err := c.Close(); err != nil {
			e.Handle(err)
		}
	}(client)
	if !initialized {
		if err := gdb.Migrate(client, ctx); err != nil {
			e.Handle(err)
			return
		}
	}

	if err := route(args[1:], client, ctx); err != nil {
		e.Handle(err)
		return
	}
}

func route(args []string, gcl *gent.Client, ctx context.Context) (ret error) {
	if len(args) == 0 {
		return failure.New(e.Operation, failure.Message("subcommand must be specified"))
	}

	// init is exceptional
	if args[0] == "init" {
		return failure.Wrap(initRoute(ctx))
	}

	isInitialized, err := directories.Exists(skeleton.Path)
	if err != nil {
		return failure.Wrap(err)
	}
	if !isInitialized {
		return failure.New(e.Operation, failure.Message("not a tit repository"))
	}

	// Prepare Database
	lcl, err := local.MakeClient(local.FilePath)
	if err != nil {
		return failure.Wrap(err)
	}
	defer func(client *lent.Client) {
		ret = multierr.Append(ret, failure.Translate(client.Close(), e.Database))
	}(lcl)
	if err := local.Migrate(lcl, ctx); err != nil {
		e.Handle(err)
		return
	}

	// Routing
	switch args[0] {
	case "remote":
		return failure.Wrap(remoteRoute(args[1:], lcl, ctx))
	case "login":
		return failure.Wrap(loginRoute(args[1:], gcl, ctx))
	case "add":
		return failure.Wrap(addRoute(args[1:], lcl, ctx))
	case "status":
		return failure.Wrap(statusRoute(lcl, ctx))
	case "commit":
		return failure.Wrap(commitRoute(args[1:], lcl, ctx))
	case "push":
		return failure.Wrap(pushRoute(args[1:], gcl, lcl, ctx))
	default:
		return failure.New(e.Operation, failure.Messagef("subcommand: \"%s\" does not exits", args[0]))
	}
}

func initRoute(ctx context.Context) (err error) {
	return failure.Wrap(subcommands.Init(ctx))
}

func remoteRoute(args []string, client *lent.Client, ctx context.Context) (err error) {
	if len(args) > 0 {
		switch args[0] {
		case "add":
			return failure.Wrap(remoteAddRoute(args[1:], client, ctx))
		}
	}
	return failure.Wrap(subcommands.Remote(args, os.Stdout, client, ctx))
}

func remoteAddRoute(args []string, client *lent.Client, ctx context.Context) (err error) {
	return failure.Wrap(remote.Add(args, client, ctx))
}

func loginRoute(args []string, cl *gent.Client, ctx context.Context) error {
	return failure.Wrap(subcommands.Login(args, os.Stdin, os.Stdout, cl, ctx))
}

func addRoute(args []string, cl *lent.Client, ctx context.Context) error {
	return failure.Wrap(subcommands.Add(args, cl, ctx))
}

func statusRoute(cl *lent.Client, ctx context.Context) error {
	return failure.Wrap(subcommands.Status(os.Stdout, cl, ctx))
}

func commitRoute(args []string, cl *lent.Client, ctx context.Context) error {
	return failure.Wrap(subcommands.Commit(args, cl, ctx))
}

func pushRoute(args []string, gcl *gent.Client, lcl *lent.Client, ctx context.Context) error {
	return failure.Wrap(subcommands.Push(args, gcl, lcl, ctx))
}
