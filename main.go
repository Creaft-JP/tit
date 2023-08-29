package main

import (
	"bytes"
	"context"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/subcommands"
	"github.com/Creaft-JP/tit/subcommands/remote"
	"github.com/Creaft-JP/tit/types"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"os"
)

func main() {
	args := os.Args
	ctx := context.Background()
	if err := route(args[1:], ctx); err != nil {
		e.Handle(err)
		os.Exit(1)
	}
}

func route(args []string, ctx context.Context) error {
	if len(args) > 0 {
		switch args[0] {
		case "init":
			return initRoute(ctx)
		case "remote":
			return remoteRoute(args[1:])
		}
	}
	panic("Not Found")
}

func initRoute(ctx context.Context) (err error) {
	return failure.Wrap(subcommands.Init(ctx))
}

func remoteRoute(args []string) (err error) {
	if len(args) > 0 {
		switch args[0] {
		case "add":
			return remoteAddRoute(args[1:])
		}
	}
	configReader, err := os.Open(types.ConfigFilepath)
	if err != nil {
		return failure.Translate(err, e.File)
	}
	defer func() {
		err = multierr.Append(err, failure.Translate(configReader.Close(), e.File))
	}()
	return subcommands.Remote(args, configReader, os.Stdout)
}

func remoteAddRoute(args []string) (err error) {
	configReader, err := os.Open(types.ConfigFilepath)
	if err != nil {
		return failure.Translate(err, e.File)
	}
	defer func() {
		err = multierr.Append(err, failure.Translate(configReader.Close(), e.File))
	}()
	configWriter := bytes.NewBuffer([]byte{})
	if err := remote.Add(args, configReader, configWriter); err != nil {
		return failure.Wrap(err)
	}
	configFile, err := os.Create(types.ConfigFilepath)
	if err != nil {
		return failure.Translate(err, e.File)
	}
	defer func() {
		err = multierr.Append(err, failure.Translate(configFile.Close(), e.File))
	}()
	if _, err := configFile.Write(configWriter.Bytes()); err != nil {
		return failure.Translate(err, e.File)
	}
	return nil
}
