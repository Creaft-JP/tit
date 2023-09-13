package test

import (
	"bytes"
	"context"
	"entgo.io/ent/dialect"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/enttest"
	"os"
	"path/filepath"
	"testing"
)

var prevd string

// SetUp returns ent.Client, io.Writer for console and context.Context
func SetUp(t *testing.T) (*ent.Client, *bytes.Buffer, context.Context) {
	currd, err := filepath.Abs(".")
	prevd = currd
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir("testdata", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir("testdata"); err != nil {
		t.Fatal(err)
	}
	return enttest.Open(t, dialect.SQLite, "./test_db?_fk=1"), bytes.NewBuffer([]byte{}), context.Background()
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
