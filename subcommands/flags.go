package subcommands

import (
	"flag"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
)

func parse(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return failure.Translate(err, e.Operation, failure.Message("failed to parse options"))
	}
	return nil
}
