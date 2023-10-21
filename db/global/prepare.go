package global

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/Creaft-JP/tit/db"
	"github.com/Creaft-JP/tit/db/global/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/global"
	"github.com/morikuni/failure"
	"path/filepath"
)

var FilePath = filepath.Join(global.Path, "database")

// MakeClient make SQLite3 ent Client from path
//
// if there isn't file on the filePath, automatically make new file
func MakeClient(filePath string) (*ent.Client, error) {
	name := fmt.Sprintf("%s?%s", filePath, db.Parameters)
	client, err := ent.Open(dialect.SQLite, name)
	return client, failure.Translate(err, e.Database)
}

func Migrate(client *ent.Client, ctx context.Context) error {
	if err := client.Schema.Create(ctx); err != nil {
		return failure.Translate(err, e.Database)
	}
	return nil
}
