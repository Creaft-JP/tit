package sqlite

import (
	"database/sql"
	"entgo.io/ent/dialect"
	"fmt"
	cdb "github.com/Creaft-JP/tit/db"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	_ "modernc.org/sqlite"
)

func Register(f string) (err error) {
	for _, driver := range sql.Drivers() {
		if driver == dialect.SQLite {
			return nil
		}
	}
	db, err := sql.Open("sqlite", fmt.Sprintf("%s?%s", f, cdb.Parameters))
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	defer func(d *sql.DB) {
		err = multierr.Append(err, failure.Translate(d.Close(), e.Database))
	}(db)
	sql.Register(dialect.SQLite, db.Driver())
	return nil
}
