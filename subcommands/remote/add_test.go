package remote

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/Creaft-JP/tit/ent"
	"github.com/Creaft-JP/tit/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func setUp(t *testing.T) *ent.Client {
	return enttest.Open(t, dialect.SQLite, "./test_db?_fk=1")
}

func tearDown(t *testing.T, client *ent.Client) {
	if err := client.Close(); err != nil {
		t.Fatalf("failed to close client: %s", err.Error())
	}
	if err := os.Remove("./test_db"); err != nil {
		t.Fatalf("failed to remove: %s", err.Error())
	}
}

func TestFirstRemoteRegister(t *testing.T) {
	// Arrange
	client := setUp(t)
	defer tearDown(t, client)

	ctx := context.Background()

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
	client := setUp(t)
	defer tearDown(t, client)

	ctx := context.Background()

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
	//client := setUp(t)
	//defer tearDown(t, client)
	//
	//ctx := context.Background()
	//
	//var args []string
	//
	//// Act
	//err := Add(args, client, ctx)
	//
	//// Assert
	//if err == nil {
	//	t.Fatal("An Error should be thrown, but was not.")
	//	return
	//}
	//setUp()
	//args := []string{"origin1", "https://api.tithub.tech/creaft/repo1"}
	//
	//if err := Add(args, reader, writer); err != nil {
	//	t.Errorf(err.Error())
	//	return
	//}
	//
	//var configJson types.Config
	//if err := json.Unmarshal(writer.Bytes(), &configJson); err != nil {
	//	t.Errorf("JSON parse failed.  %s", err.Error())
	//	return
	//}
	//
	//want := []config.Remote{
	//	{"origin", "https://api.tithub.tech/creaft/repository"},
	//	{"origin1", "https://api.tithub.tech/creaft/repo1"},
	//}
	//
	//got := configJson.Remotes
	//if !reflect.DeepEqual(want, got) {
	//	t.Errorf("remotes shouled be %s, but got %s.", want, got)
	//}
}

func TestBlockReplace(t *testing.T) {
	//setUp()
	//args := []string{"origin", "https://api.tithub.tech/creaft/repo1"}
	//err := Add(args, reader, writer)
	//if err == nil {
	//	t.Error("an error should be thrown, but was not.")
	//	return
	//}
	//want := "remote origin already exists"
	//got, _ := failure.MessageOf(err)
	//if got != want {
	//	t.Errorf("error message should be \"%s\", but got \"%s\".", want, got)
	//}
	//if writer.Len() > 0 {
	//	t.Errorf("no bytes should be written, but \"%s\" were.", writer.String())
	//}
}
