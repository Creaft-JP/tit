package subcommands

import (
	"bytes"
	"encoding/json"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"io"
	"testing"
)

var consoleWriter *bytes.Buffer
var configReader io.Reader

func setUp() {
	configBytes, _ := json.Marshal(types.Config{Remotes: []config.Remote{
		{"origin", "https://tithub.tech/user/repository"},
		{"test", "https://tithub.tech/test/repository"},
	}})
	configReader = bytes.NewReader(configBytes)
	consoleWriter = bytes.NewBuffer([]byte{})
}

func TestNoArgs(t *testing.T) {
	setUp()
	if err := Remote([]string{}, configReader, consoleWriter); err != nil {
		t.Error(err.Error())
		return
	}
	want := "origin\ntest\n"
	got := consoleWriter.String()
	if want != got {
		t.Errorf("console output should be \"%s\", but got \"%s\".", want, got)
	}
}
func testVerbose(t *testing.T, args []string) {
	setUp()
	if err := Remote(args, configReader, consoleWriter); err != nil {
		t.Error(err.Error())
		return
	}
	want := `origin https://tithub.tech/user/repository
test https://tithub.tech/test/repository
`
	got := consoleWriter.String()
	if want != got {
		t.Errorf("console output should be \"%s\", but got \"%s\".", want, got)
	}
}
func TestVerbose(t *testing.T) {
	testVerbose(t, []string{"--verbose"})
}
func TestV(t *testing.T) {
	testVerbose(t, []string{"-v"})
}
