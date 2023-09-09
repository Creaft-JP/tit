package subcommands

import (
	"bufio"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/Creaft-JP/tit/db/global/ent"
	"github.com/Creaft-JP/tit/db/global/ent/logintoken"
	e "github.com/Creaft-JP/tit/error"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"golang.org/x/term"
	"io"
	"os"
)

func Login(args []string, r io.Reader, w io.Writer, cl *ent.Client, ctx context.Context) error {
	fs := flag.NewFlagSet("login", flag.ContinueOnError)
	slug := fs.String("u", "", "user slug")
	if err := parse(fs, args); err != nil {
		return failure.Wrap(err)
	}
	if *slug == "" {
		if _, err := fmt.Fprint(w, "user slug: "); err != nil {
			return failure.Translate(err, e.File)
		}
		scanner := bufio.NewScanner(r)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return failure.Translate(err, e.File)
			}
		}
		*slug = scanner.Text()
	}
	if _, err := fmt.Fprint(w, "token: "); err != nil {
		return failure.Translate(err, e.File)
	}
	token, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return failure.Translate(err, e.File)
	}
	transaction, err := cl.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return failure.Translate(err, e.Database)
	}
	if _, err := transaction.LoginToken.Delete().Where(logintoken.SignInUserSlug(*slug)).Exec(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), failure.Translate(transaction.Rollback(), e.Database))
	}
	if _, err := transaction.LoginToken.Create().SetSignInUserSlug(*slug).SetCliLoginToken(string(token)).Save(ctx); err != nil {
		return multierr.Append(failure.Translate(err, e.Database), failure.Translate(transaction.Rollback(), e.Database))
	}
	return failure.Translate(transaction.Commit(), e.Database)
}
