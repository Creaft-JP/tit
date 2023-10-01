package subcommands

import (
	"context"
	"fmt"
	"github.com/Creaft-JP/tit/db/local/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"io"
)

func Status(w io.Writer, cl *ent.Client, ctx context.Context) error {
	files, err := cl.StagedFile.Query().All(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	for _, file := range files {
		if _, err := fmt.Fprintln(w, file.Path); err != nil {
			return failure.Translate(err, e.File)
		}
	}
	return nil
}
