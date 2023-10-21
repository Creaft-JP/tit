package remote

import (
	"github.com/Creaft-JP/tit/db/local/ent/remote"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/test"
	"github.com/morikuni/failure"
	"testing"
)

func TestFirstRemoteRegister(t *testing.T) {
	// Arrange
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)

	args := []string{"origin", "https://api.tithub.tech/creaft/repository"}

	// Act
	if err := Add(args, client, ctx); err != nil {
		t.Fatalf("failed to Add: %s", err.Error())
	}

	// Assert
	got, err := client.Remote.Query().All(ctx)
	if err != nil {
		t.Fatalf("failed to get remotes: %s", err.Error())
	}
	if len(got) != 1 {
		t.Fatalf("remotes count should be 1, but got: %d", len(got))
	}
	if got[0].Name != "origin" {
		t.Fatalf("remote Name should be %s, but got: %s", "origin", got[0].Name)
	}
	if got[0].URL != "https://api.tithub.tech/creaft/repository" {
		t.Fatalf("remote URL should be %s, but got: %s", "https://api.tithub.tech/creaft/repository", got[0].URL)
	}
}

func TestLackOfArgumentsError(t *testing.T) {
	// Arrange
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)

	var args []string

	// Act
	err := Add(args, client, ctx)

	// Assert
	if err == nil {
		t.Fatal("An Error should be thrown, but was not.")
		return
	}
}

func TestSecondRemoteRegister(t *testing.T) {
	// Arrange
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)

	if err := Add([]string{"origin", "https://api.tithub.tech/creaft/repo"}, client, ctx); err != nil {
		t.Fatalf("failed to Add first: %s", err.Error())
	}

	// Act
	if err := Add([]string{"copy", "https://api.tithub.tech/creaft/copy"}, client, ctx); err != nil {
		t.Fatalf("failed to Add second: %s", err.Error())
	}

	// Assert
	got, err := client.Remote.Query().Order(remote.ByName()).All(ctx)
	if err != nil {
		t.Fatalf("failed to get remotes: %s", err.Error())
	}
	if len(got) != 2 {
		t.Fatalf("remotes count should be 2, but got: %d", len(got))
	}
	wantNames := []string{"copy", "origin"}
	for i, gotRemote := range got {
		if gotRemote.Name != wantNames[i] {
			t.Fatalf("remote Name should be %s, but got: %s", wantNames[i], gotRemote.Name)
		}
	}
	wantUrls := []string{"https://api.tithub.tech/creaft/copy", "https://api.tithub.tech/creaft/repo"}
	for i, gotRemote := range got {
		if gotRemote.URL != wantUrls[i] {
			t.Fatalf("remote URL should be %s, but got: %s", wantUrls[i], gotRemote.URL)
		}
	}
}

func TestBlockReplace(t *testing.T) {
	// Arrange
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)

	if err := Add([]string{"origin", "https://api.tithub.tech/creaft/repo"}, client, ctx); err != nil {
		t.Fatalf("failed to Add first: %s", err.Error())
	}

	// Act
	err := Add([]string{"origin", "https://api.tithub.tech/creaft/copy"}, client, ctx)

	// Assert
	if err == nil {
		t.Fatal("an error should be thrown, but was not.")
	}

	code, _ := failure.CodeOf(err)
	if code != e.Operation {
		t.Fatalf("error code should be %s, but got %s", e.Operation, code)
	}

	message, _ := failure.MessageOf(err)
	wantMessage := "remote origin already exists"
	if message != wantMessage {
		t.Fatalf("error message should be \"%s\", but got \"%s\".", wantMessage, message)
	}

	gotRemote, err := client.Remote.Query().Where(remote.Name("origin")).Only(ctx)
	if err != nil {
		t.Fatalf("failed to get remote: %s", err.Error())
	}
	wantUrl := "https://api.tithub.tech/creaft/repo"
	if gotRemote.URL != wantUrl {
		t.Fatalf("remote URL should be %s, but got: %s", wantUrl, gotRemote.URL)
	}
}
