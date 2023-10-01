package subcommands

import (
	"context"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/test"
	"testing"
)

func setUpRemoteTest(t *testing.T, client *ent.Client, ctx context.Context) {
	_, err := client.Remote.CreateBulk(
		client.Remote.Create().SetName("origin").SetURL("https://tithub.tech/user/repository"),
		client.Remote.Create().SetName("test").SetURL("https://tithub.tech/test/repository"),
	).Save(ctx)
	if err != nil {
		t.Fatalf("failed to bulk insert remotes: %s", err.Error())
	}
}

func TestNoArgs(t *testing.T) {
	// Arrange
	client, writer, ctx := test.SetUp(t)
	defer test.TearDown(t, client)

	setUpRemoteTest(t, client, ctx)

	// Act
	if err := Remote([]string{}, writer, client, ctx); err != nil {
		t.Fatalf("failed to act: %s", err.Error())
		return
	}

	// Assert
	want := "origin\ntest\n"
	got := writer.String()
	if want != got {
		t.Errorf("console output should be \"%s\", but got \"%s\".", want, got)
	}
}

func testVerbose(t *testing.T, args []string) {
	// Arrange
	client, writer, ctx := test.SetUp(t)
	defer test.TearDown(t, client)

	setUpRemoteTest(t, client, ctx)

	// Act
	if err := Remote(args, writer, client, ctx); err != nil {
		t.Error(err.Error())
		return
	}

	// Assert
	want := `origin https://tithub.tech/user/repository
test https://tithub.tech/test/repository
`
	got := writer.String()
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
