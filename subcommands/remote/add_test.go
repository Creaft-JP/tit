package remote

import (
	"bytes"
	"encoding/json"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"github.com/google/go-cmp/cmp"
	"io"
	"testing"
)

var emptyInitialConfig, _ = json.Marshal(types.Config{
	Remotes: []config.Remote{},
})
var initialConfig, _ = json.Marshal(types.Config{Remotes: []config.Remote{{"origin", "https://api.tithub.tech/creaft/repository"}}})

var emptyReader io.Reader
var reader io.Reader
var writer *bytes.Buffer

func setUp() {
	writer = &bytes.Buffer{}
	emptyReader = bytes.NewReader(emptyInitialConfig)
	reader = bytes.NewReader(initialConfig)
}

func TestFirstRemoteRegister(t *testing.T) {
	setUp()
	args := []string{"origin", "https://api.tithub.tech/creaft/repository"}
	if err := Add(args, emptyReader, writer); err != nil {
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

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("got is not equals got, diff: \n%s", diff)
	}
}

func TestLackOfArgumentsError(t *testing.T) {
	setUp()
	args := []string{"origin"}
	err := Add(args, emptyReader, writer)
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
	setUp()
	args := []string{"origin", "https://api.tithub.tech/creaft/repository", "ssh"}
	err := Add(args, emptyReader, writer)
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
	setUp()
	args := []string{"origin1", "https://api.tithub.tech/creaft/repo1"}

	initialConfig, err := json.Marshal(types.Config{
		Remotes: []config.Remote{
			{"origin", "https://api.tithub.tech/creaft/repository"},
		},
	})
	if err != nil {
		t.Error(err.Error())
	}

	emptyReader = bytes.NewReader(initialConfig)
	if err := Add(args, emptyReader, writer); err != nil {
		t.Errorf(err.Error())
	}

	var configJson types.Config
	if err := json.Unmarshal(writer.Bytes(), &configJson); err != nil {
		t.Errorf(`
JSON parse Failed
Reason: %s
Got:
%s
`, err.Error(), writer.String())
		return
	}

	want := []config.Remote{
		{"origin", "https://api.tithub.tech/creaft/repository"},
		{"origin1", "https://api.tithub.tech/creaft/repo1"},
	}

	got := configJson.Remotes
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("got is not equals got, diff: \n%s", diff)
	}
}

func TestBlockReplace(t *testing.T) {
	setUp()
	args := []string{"origin", "https://api.tithub.tech/creaft/repo1"}
	err := Add(args, reader, writer)
	if err == nil {
		t.Error("an error should be thrown, but was not.")
	}
	want := "remote origin already exists"
	got := err.Error()
	if got != want {
		t.Errorf("error message should be \"%s\", but got \"%s\".", want, got)
	}
	if writer.Len() > 0 {
		t.Errorf("no bytes should be written, but \"%s\" were.", writer.String())
	}
}
