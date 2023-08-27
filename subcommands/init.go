package subcommands

import (
	"encoding/json"
	"fmt"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"github.com/morikuni/failure"
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
	if err := os.Mkdir(name, perm); err != nil {
		return failure.Translate(err, e.File)
	}
	return nil
}
func (_ *osFileDirectoryCreator) create(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, failure.Translate(err, e.File)
	}
	return file, nil
}

type fileDirectoryStatusReader interface {
	stat(name string) (os.FileInfo, error)
}

type osFileDirectoryStatusReader struct{}

func (_ *osFileDirectoryStatusReader) stat(name string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(name)
	if err == nil {
		return fileInfo, nil
	}
	if os.IsNotExist(err) {
		return nil, failure.Translate(err, e.FileNotFound)
	}
	return nil, failure.Translate(err, e.File)
}

func CreateRepository() error {
	if err := checkAlreadyInitialized(&osFileDirectoryStatusReader{}); err != nil {
		return failure.Wrap(err)
	}
	if err := createFiles(&osFileDirectoryCreator{}); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
func Init(consoleWriter io.Writer, configWriter io.Writer) error {
	if err := initConfig(configWriter); err != nil {
		return failure.Wrap(err)
	}
	if err := initMessage(consoleWriter); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
func initConfig(writer io.Writer) error {
	if err := json.NewEncoder(writer).Encode(types.Config{Remotes: []config.Remote{}}); err != nil {
		return failure.Translate(err, e.File)
	}
	return nil
}
func initMessage(writer io.Writer) error {
	if _, err := fmt.Fprintf(writer, "Initialized empty Tit repository in ./%s/\n", types.RepositoryDirectoryName); err != nil {
		return failure.Translate(err, e.File)
	}
	return nil
}
func checkAlreadyInitialized(reader fileDirectoryStatusReader) error {
	_, err := reader.stat(types.RepositoryDirectoryName)
	if err == nil {
		return failure.New(e.Operation, failure.Message("tit repository already exists"))
	}
	if code, _ := failure.CodeOf(err); code == e.FileNotFound {
		return nil
	}
	return failure.Wrap(err)
}
func createFiles(creator fileDirectoryCreator) error {
	if err := creator.mkdir(types.RepositoryDirectoryName, os.FileMode(0755)); err != nil {
		return failure.Wrap(err)
	}
	if _, err := creator.create(types.ConfigFilepath); err != nil {
		return failure.Wrap(err)
	}
	return nil
}
