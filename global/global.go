package global

import (
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"os"
	"path/filepath"
)

var Path = ""

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		e.Handle(failure.Translate(err, e.File))
		return
	}
	Path = filepath.Join(home, ".tit-global")
}
