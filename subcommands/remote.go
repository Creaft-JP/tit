package subcommands

import (
	"context"
	"flag"
	"fmt"
	"github.com/Creaft-JP/tit/db/local/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"io"
)

func Remote(args []string, consoleWriter io.Writer, client *ent.Client, ctx context.Context) error {
	fs := flag.NewFlagSet("remote", flag.ContinueOnError)
	var verbose bool
	setVerbose(fs, &verbose, "verbose")
	setVerbose(fs, &verbose, "v")
	if err := parse(fs, args); err != nil {
		return failure.Wrap(err)
	}
	remotes, err := client.Remote.Query().All(ctx)
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	for _, remote := range remotes {
		if err := printRemote(consoleWriter, remote, verbose); err != nil {
			return failure.Wrap(err)
		}
	}
	return nil
}

func setVerbose(flagSet *flag.FlagSet, verbose *bool, name string) {
	flagSet.BoolVar(verbose, name, false, "print a detailed information about remotes")
}

func printRemote(writer io.Writer, remote *ent.Remote, verbose bool) error {
	if verbose {
		if _, err := fmt.Fprintf(writer, "%s %s\n", remote.Name, remote.URL); err != nil {
			return failure.Translate(err, e.File)
		}
	} else {
		if _, err := fmt.Fprintln(writer, remote.Name); err != nil {
			return failure.Translate(err, e.File)
		}
	}
	return nil
}
