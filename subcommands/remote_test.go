package subcommands

import (
	"bytes"
	"encoding/json"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"io"
	"testing"
)

type remoteTestingValue struct {
	configReader  io.Reader
	consoleWriter *bytes.Buffer
}

func setUpRemoteTest() remoteTestingValue {
	configBytes, _ := json.Marshal(types.Config{Remotes: []config.Remote{
		{"origin", "https://tithub.tech/user/repository"},
		{"test", "https://tithub.tech/test/repository"},
	}})
	return remoteTestingValue{
		configReader:  bytes.NewReader(configBytes),
		consoleWriter: bytes.NewBuffer([]byte{}),
	}
}

func TestNoArgs(t *testing.T) {
	tv := setUpRemoteTest()
	if err := Remote([]string{}, tv.configReader, tv.consoleWriter); err != nil {
		t.Error(err.Error())
		return
	}
	want := "origin\ntest\n"
	got := tv.consoleWriter.String()
	if want != got {
		t.Errorf("console output should be \"%s\", but got \"%s\".", want, got)
	}
}
func testVerbose(t *testing.T, args []string) {
	tv := setUpRemoteTest()
	if err := Remote(args, tv.configReader, tv.consoleWriter); err != nil {
		t.Error(err.Error())
		return
	}
	want := `origin https://tithub.tech/user/repository
test https://tithub.tech/test/repository
`
	got := tv.consoleWriter.String()
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
