package test

import (
	"bytes"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/Creaft-JP/tit/db"
	"github.com/Creaft-JP/tit/db/local"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/enttest"
	"github.com/Creaft-JP/tit/sqlite"
	"os"
	"path/filepath"
	"testing"
)

var prevd string
var dsn = fmt.Sprintf("%s?%s", local.FilePath, db.Parameters)

// SetUp returns ent.Client, io.Writer for console and context.Context
func SetUp(t *testing.T) (*ent.Client, *bytes.Buffer, context.Context) {
	currd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	prevd = currd
	if err := os.Mkdir("testdata", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir("testdata"); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(".tit", 0755); err != nil {
		t.Fatal(err)
	}
	if err := sqlite.Register(dsn); err != nil {
		t.Fatal(err)
	}
	return enttest.Open(t, dialect.SQLite, dsn), bytes.NewBuffer([]byte{}), context.Background()
}
func SetUpInner(t *testing.T) (*ent.Client, context.Context) {
	currd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	prevd = currd
	newd := filepath.Join("testdata", "repository")
	if err := os.MkdirAll(newd, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(newd); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(".tit", 0755); err != nil {
		t.Fatal(err)
	}
	return enttest.Open(t, dialect.SQLite, dsn), context.Background()
}

func TearDown(t *testing.T, client *ent.Client) {
	if err := client.Close(); err != nil {
		t.Fatalf("failed to close client: %s", err.Error())
	}
	if err := os.Chdir(prevd); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll("testdata"); err != nil {
		t.Fatalf("failed to remove: %s", err.Error())
	}
}
