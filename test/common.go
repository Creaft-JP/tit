package test

import (
	"bytes"
	"entgo.io/ent/dialect"
	"github.com/Creaft-JP/tit/ent"
	"github.com/Creaft-JP/tit/ent/enttest"
	"os"
	"testing"
)

// SetUp return ent.Client and io.Writer for console
func SetUp(t *testing.T) (*ent.Client, *bytes.Buffer) {
	return enttest.Open(t, dialect.SQLite, "./test_db?_fk=1"), bytes.NewBuffer([]byte{})
}

func TearDown(t *testing.T, client *ent.Client) {
	if err := client.Close(); err != nil {
		t.Fatalf("failed to close client: %s", err.Error())
	}
	if err := os.Remove("./test_db"); err != nil {
		t.Fatalf("failed to remove: %s", err.Error())
	}
}
