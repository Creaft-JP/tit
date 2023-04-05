package remote

import (
	"bytes"
	"encoding/json"
	"github.com/Creaft-JP/tit/types/config"
	"reflect"
	"testing"
)

import (
	"github.com/Creaft-JP/tit/types"
)

var initialConfig = types.Config{
	Remotes: []config.Remote{},
}

var reader *bytes.Buffer
var writer *bytes.Buffer

func TestMain(m *testing.M) {
	reader = &bytes.Buffer{}
	writer = &bytes.Buffer{}
	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(initialConfig); err != nil {
		panic(err)
		return
	}
	m.Run()
}

func TestFirstRemoteRegister(t *testing.T) {
	args := []string{"origin", "https://api.tithub.tech/creaft/repository"}
	if err := Add(args, reader, writer); err != nil {
		t.Errorf(err.Error())
	}

	var configJson types.Config
	err := json.Unmarshal(writer.Bytes(), &configJson)
	if err != nil {
		t.Errorf("JSON parse failed, %s", err.Error())
		return
	}

	want := []config.Remote{
		{Name: "origin", Url: "https://api.tithub.tech/creaft/repository"},
	}
	got := configJson.Remotes

	if reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestLackOfArgumentsError(t *testing.T) {
	args := []string{"origin"}
	err := Add(args, reader, writer)
	if err == nil {
		t.Error("An Error should be thrown, but was not.")
	}
	want := "args should be 2, but received 1"
	got := err.Error()
	if got != want {
		t.Errorf("thrown error message should be \"%s\", but got \"%s\"", want, got)
	}
}
func TestTooManyArgumentsError(t *testing.T) {
	args := []string{"origin", "https://api.tithub.tech/creaft/repository", "ssh"}
	err := Add(args, reader, writer)
	if err == nil {
		t.Error("An Error should be thrown, but was not.")
	}
	want := "args should be 2, but received 3"
	got := err.Error()
	if got != want {
		t.Errorf("thrown error message should be \"%s\", but got \"%s\"", want, got)
	}
}

func TestSecondRemoteRegister(t *testing.T) {
	args := []string{"origin1", "https://api.tithub.tech/creaft/repo1"}
	encoder := json.NewEncoder(reader)
	initialConfig := types.Config{
		Remotes: []config.Remote{
			{"origin", "https://api.tithub.tech/creaft/repository"},
		},
	}
	if err := encoder.Encode(initialConfig); err != nil {
		t.Errorf("initial config writing has thrown Error: %v", err.Error())
		return
	}
	if err := Add(args, reader, writer); err != nil {
		t.Errorf(err.Error())
	}

	var configJson types.Config
	err := json.Unmarshal(writer.Bytes(), &configJson)
	if err != nil {
		t.Errorf("JSON parse failed, %s", err.Error())
		return
	}

	want := []config.Remote{
		{"origin", "https://api.tithub.tech/creaft/repository"},
		{"origin1", "https://api.tithub.tech/creaft/repo1"},
	}

	got := configJson.Remotes

	if !reflect.DeepEqual(got, want) {
		t.Errorf("remotes should be %v, but was written %v", want, got)
	}
}
