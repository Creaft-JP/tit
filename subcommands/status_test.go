package subcommands

import (
	"github.com/Creaft-JP/tit/test"
	"testing"
)

func TestStatus(t *testing.T) {
	client, writer, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	if _, err := client.StagedFile.CreateBulk(
		client.StagedFile.Create().SetPath("new-file-1").SetContent("Hello, world!!"),
		client.StagedFile.Create().SetPath("new-file-2").SetContent("Hello, tit!!"),
	).Save(ctx); err != nil {
		t.Fatal(err)
	}
	if err := Status(writer, client, ctx); err != nil {
		t.Fatal(err)
	}
	want := "new-file-1\nnew-file-2\n"
	got := writer.String()
	if got != want {
		t.Errorf("output should be \"%s\", but got \"%s\"", want, got)
	}
}
