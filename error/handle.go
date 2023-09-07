package error

import (
	"fmt"
	"github.com/morikuni/failure"
	"os"
)

func Handle(err error) {
	if err == nil {
		return
	}

	code, _ := failure.CodeOf(err)
	if code == Operation {
		message, _ := failure.MessageOf(err)
		if _, err := fmt.Fprintln(os.Stderr, message); err != nil {
			panic(err)
		}
		return
	}

	if _, e := fmt.Fprintln(os.Stderr, "============ Error ============"); e != nil {
		panic(e)
	}
	if _, e := fmt.Fprintf(os.Stderr, "Error = %v\n", err); e != nil {
		panic(e)
	}

	if _, e := fmt.Fprintf(os.Stderr, "Code = %v\n", code); e != nil {
		panic(e)
	}

	msg, _ := failure.MessageOf(err)
	if _, e := fmt.Fprintf(os.Stderr, "Message = %v\n", msg); e != nil {
		panic(e)
	}

	cs, _ := failure.CallStackOf(err)
	if _, e := fmt.Fprintf(os.Stderr, "CallStack = %v\n", cs); e != nil {
		panic(e)
	}

	if _, e := fmt.Fprintf(os.Stderr, "Cause = %v\n", failure.CauseOf(err)); e != nil {
		panic(e)
	}

	if _, e := fmt.Fprintln(os.Stderr); e != nil {
		panic(e)
	}
	if _, e := fmt.Fprintln(os.Stderr, "============ Detail ============"); e != nil {
		panic(e)
	}
	if _, e := fmt.Fprintf(os.Stderr, "%+v\n", err); e != nil {
		panic(e)
	}
}
