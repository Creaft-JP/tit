package subcommands

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"os"
	"reflect"
	"testing"
)

type initTestingValue struct {
	consoleWriter *bytes.Buffer
	configWriter  *bytes.Buffer
}

func setUpInitTest() initTestingValue {
	return initTestingValue{
		consoleWriter: bytes.NewBuffer([]byte{}),
		configWriter:  bytes.NewBuffer([]byte{}),
	}
}

type mockFileDirectoryCreator struct {
	directoryName string
	perm          os.FileMode
	filename      string
}

func (creator *mockFileDirectoryCreator) mkdir(name string, perm os.FileMode) error {
	creator.directoryName = name
	creator.perm = perm
	return nil
}
func (creator *mockFileDirectoryCreator) create(filename string) (*os.File, error) {
	creator.filename = filename
	return nil, nil
}
func TestMakeDirectory(t *testing.T) {
	mock := &mockFileDirectoryCreator{}
	if err := createFiles(mock); err != nil {
		t.Error(err.Error())
		return
	}
	wantDirectoryName := ".tit"
	gotDirectoryName := mock.directoryName
	if wantDirectoryName != gotDirectoryName {
		t.Errorf("directory name should be \"%s\", but got \"%s\".", wantDirectoryName, gotDirectoryName)
	}
	wantPerm := os.FileMode(0755)
	gotPerm := mock.perm
	if wantPerm != gotPerm {
		t.Errorf("permission should be \"%s\", but got \"%s\".", wantPerm, gotPerm)
	}
	wantFilename := ".tit/config.json"
	gotFilename := mock.filename
	if wantFilename != gotFilename {
		t.Errorf("filename should be \"%s\", but got \"%s\".", wantFilename, gotFilename)
	}
}
func TestInitConfig(t *testing.T) {
	tv := setUpInitTest()
	if err := initConfig(tv.configWriter); err != nil {
		t.Error(err.Error())
		return
	}
	want := types.Config{Remotes: []config.Remote{}}
	decoder := json.NewDecoder(tv.configWriter)
	var got types.Config
	if err := decoder.Decode(&got); err != nil {
		t.Error(err.Error())
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("config should be %s, but got %s.", want, got)
	}
}
func TestInitMessage(t *testing.T) {
	tv := setUpInitTest()
	if err := initMessage(tv.consoleWriter); err != nil {
		t.Error(err.Error())
		return
	}

	want := "Initialized empty Tit repository in ./.tit/\n"
	got := tv.consoleWriter.String()
	if want != got {
		t.Errorf("console output should be \"%s\", but got \"%s\".", want, got)
	}
}

type mockFileDirectoryStatusReader struct {
	err  error
	name string
}

func (reader *mockFileDirectoryStatusReader) stat(name string) (os.FileInfo, error) {
	reader.name = name
	return nil, reader.err
}
func TestCheckAlreadyInitialized(t *testing.T) {
	mock := &mockFileDirectoryStatusReader{nil, ""}
	err := checkAlreadyInitialized(mock)
	if err == nil {
		t.Errorf("err should be not nil")
		return
	}
	want := "tit repository already exists"
	got := err.Error()
	if want != got {
		t.Errorf("error should be \"%s\", but got \"%s\".", want, got)
	}

	mock.assertStatRequestName(t)
}
func TestNotAlreadyInitialized(t *testing.T) {
	mock := &mockFileDirectoryStatusReader{os.ErrNotExist, ""}
	if err := checkAlreadyInitialized(mock); err != nil {
		t.Error(err)
		return
	}

	mock.assertStatRequestName(t)
}
func TestOtherErrorByStat(t *testing.T) {
	otherError := errors.New("other error")
	mock := &mockFileDirectoryStatusReader{otherError, ""}
	err := checkAlreadyInitialized(mock)
	if err == nil {
		t.Error("Error should be thrown, but not.")
		return
	}
	want := otherError
	got := err
	if want != got {
		t.Errorf("Error should be %s, but got %s.", want, got)
	}
}
func (reader *mockFileDirectoryStatusReader) assertStatRequestName(t *testing.T) {
	want := ".tit"
	got := reader.name
	if want != got {
		t.Errorf("name should be \"%s\", but got \"%s\".", want, got)
	}
}
