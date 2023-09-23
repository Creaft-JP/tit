package subcommands

import (
	"context"
	"flag"
	"fmt"
	"github.com/Creaft-JP/tit/db/local/ent"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/skeleton"
	"github.com/morikuni/failure"
	"os"
	"os/exec"
	"path/filepath"
)

const commitMessageDefault = "_tit_commit_message_default_value"

func Commit(args []string, cl *ent.Client, ctx context.Context) error {
	fs := flag.NewFlagSet("commit", flag.ContinueOnError)
	message := fs.String("m", commitMessageDefault, "commit message")
	if err := parse(fs, args); err != nil {
		return failure.Wrap(err)
	}
	if *message == commitMessageDefault {
		editor, ok := os.LookupEnv("EDITOR")
		if !ok {
			editor = "vi"
		}
		tmp := filepath.Join(skeleton.Path, fmt.Sprintf("commit message (%d)", os.Getpid()))
		command := exec.Command(editor, tmp)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			return failure.Translate(err, e.Editor)
		}
		content, err := os.ReadFile(tmp)
		if err != nil {
			return failure.Translate(err, e.File)
		}
		*message = string(content)
		if err := os.Remove(tmp); err != nil {
			return failure.Translate(err, e.File)
		}
	}
	return failure.Wrap(commit(*message, cl, ctx))
}
func commit(me string, cl *ent.Client, ctx context.Context) error {
	return nil
}
