package subcommands

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"io"
	"os"
)

type fileDirectoryCreator interface {
	mkdir(name string, perm os.FileMode) error
	create(filename string) (*os.File, error)
}

type osFileDirectoryCreator struct {
}

func (_ *osFileDirectoryCreator) mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}
func (_ *osFileDirectoryCreator) create(filename string) (*os.File, error) {
	return os.Create(filename)
}

type fileDirectoryStatusReader interface {
	stat(name string) (os.FileInfo, error)
}

type osFileDirectoryStatusReader struct{}

func (_ *osFileDirectoryStatusReader) stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func CreateRepository() error {
	if err := checkAlreadyInitialized(&osFileDirectoryStatusReader{}); err != nil {
		return err
	}
	if err := createFiles(&osFileDirectoryCreator{}); err != nil {
		return err
	}
	return nil
		return err
	}
	return nil
}
func initConfig(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(types.Config{Remotes: []config.Remote{}})
}
func initMessage(writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "Initialized empty Tit repository in ./%s/\n", types.RepositoryDirectoryName)
	return err
}
func checkAlreadyInitialized(reader fileDirectoryStatusReader) error {
	_, err := reader.stat(types.RepositoryDirectoryName)
	if err == nil {
		return errors.New("tit repository already exists")
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
func createFiles(creator fileDirectoryCreator) error {
	if err := creator.mkdir(types.RepositoryDirectoryName, os.FileMode(0755)); err != nil {
		return err
	}
	if _, err := creator.create(types.ConfigFilepath); err != nil {
		return err
	}
	return nil
}
