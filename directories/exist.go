package directories

import (
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"os"
)

func Exists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, failure.Translate(err, e.File)
		}
	}
	if !info.IsDir() {
		return false, failure.New(e.Operation, failure.Messagef("the path \"%s\" is not a directory", path))
	}
	return true, nil
}
