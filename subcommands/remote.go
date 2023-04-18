package subcommands

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Creaft-JP/tit/types"
	"github.com/Creaft-JP/tit/types/config"
	"io"
)

func printRemote(writer io.Writer, remote config.Remote, verbose bool) error {
	if verbose {
		if _, err := fmt.Fprintf(writer, "%s %s\n", remote.Name, remote.Url); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintln(writer, remote.Name); err != nil {
			return err
		}
	}
	return nil
}
func Remote(args []string, configReader io.Reader, consoleWriter io.Writer) error {
	flagSet := flag.NewFlagSet("remote", flag.ContinueOnError)
	var verbose bool
	setVerbose(flagSet, &verbose, "verbose")
	setVerbose(flagSet, &verbose, "v")
	if err := flagSet.Parse(args); err != nil {
		return err
	}
	configDecoder := json.NewDecoder(configReader)
	var configContent types.Config
	if err := configDecoder.Decode(&configContent); err != nil {
		return err
	}
	for _, remote := range configContent.Remotes {
		if err := printRemote(consoleWriter, remote, verbose); err != nil {
			return err
		}
	}
	return nil
}

func setVerbose(flagSet *flag.FlagSet, verbose *bool, name string) {
	flagSet.BoolVar(verbose, name, false, "print a detailed information about remotes")
}
