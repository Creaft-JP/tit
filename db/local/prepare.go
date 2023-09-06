package local

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/Creaft-JP/tit/db/local/ent"
	e "github.com/Creaft-JP/tit/error"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morikuni/failure"
	"os"
)

const FilePath = "./.tit"

// MakeClient make SQLite3 ent Client from path
//
// if there isn't file on the filePath, automatically make new file
func MakeClient(filePath string) (*ent.Client, error) {
	name := fmt.Sprintf("%s?_fk=1", filePath)
	client, err := ent.Open(dialect.SQLite, name)
	return client, failure.Translate(err, e.Database)
}

// IsAlreadyInitialized if the filePath's file is exists, return true
func IsAlreadyInitialized(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, failure.Translate(err, e.File)
		}
	}
	if info.IsDir() {
		return false, failure.New(e.Operation, failure.Messagef("the filePath \"%s\" is directory", filePath))
	}
	return true, nil
}

func Migrate(client *ent.Client, ctx context.Context) error {
	if err := client.Schema.Create(ctx); err != nil {
		return failure.Translate(err, e.Database)
	}
	return nil
}
