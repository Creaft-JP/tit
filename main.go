package main

import (
	"bytes"
	"fmt"
	"github.com/Creaft-JP/tit/subcommands"
	"github.com/Creaft-JP/tit/subcommands/remote"
	"github.com/Creaft-JP/tit/types"
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

func initRoute() error {
	if err := subcommands.CreateRepository(); err != nil {
		return err
	}
	configWriter, err := os.Create(types.ConfigFilepath)
	if err != nil {
		return err
	}
	defer func(writer *os.File) {
		_ = writer.Close()
	}(configWriter)
	return subcommands.Init(os.Stdout, configWriter)
}

func remoteRoute(args []string) error {
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
	defer func(reader *os.File) {
		_ = reader.Close()
	}(configReader)
	return subcommands.Remote(args, configReader, os.Stdout)
}

func remoteAddRoute(args []string) error {
	configReader, err := os.Open(types.ConfigFilepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(configReader)
	configWriter := bytes.NewBuffer([]byte{})
	if err := remote.Add(args, configReader, configWriter); err != nil {
		return err
	}
	configFile, err := os.Create(types.ConfigFilepath)
	if err != nil {
		return err
	}
	defer func(writer *os.File) {
		_ = writer.Close()
	}(configFile)
	if _, err := configFile.Write(configWriter.Bytes()); err != nil {
		return err
	}
	return nil
}
