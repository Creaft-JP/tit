package main

import (
	"bytes"
	"fmt"
	"github.com/Creaft-JP/tit/subcommands"
	"github.com/Creaft-JP/tit/subcommands/remote"
	"github.com/Creaft-JP/tit/types"
	"go.uber.org/multierr"
	"os"
)

func main() {
	args := os.Args
	if err := route(args[1:]); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func route(args []string) error {
	if len(args) > 0 {
		switch args[0] {
		case "init":
			return initRoute()
		case "remote":
			return remoteRoute(args[1:])
		}
	}
	panic("Not Found")
}

func initRoute() (err error) {
	if err := subcommands.CreateRepository(); err != nil {
		return err
	}
	configWriter, err := os.Create(types.ConfigFilepath)
	if err != nil {
		return err
	}
	defer func() {
		err = multierr.Append(err, configWriter.Close())
	}()
	return subcommands.Init(os.Stdout, configWriter)
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
		return err
	}
	defer func() {
		err = multierr.Append(err, configReader.Close())
	}()
	return subcommands.Remote(args, configReader, os.Stdout)
}

func remoteAddRoute(args []string) (err error) {
	configReader, err := os.Open(types.ConfigFilepath)
	if err != nil {
		return err
	}
	defer func() {
		err = multierr.Append(err, configReader.Close())
	}()
	configWriter := bytes.NewBuffer([]byte{})
	if err := remote.Add(args, configReader, configWriter); err != nil {
		return err
	}
	configFile, err := os.Create(types.ConfigFilepath)
	if err != nil {
		return err
	}
	defer func() {
		err = multierr.Append(err, configFile.Close())
	}()
	if _, err := configFile.Write(configWriter.Bytes()); err != nil {
		return err
	}
	return nil
}
