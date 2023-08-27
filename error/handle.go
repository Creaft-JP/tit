package error

import (
	"fmt"
	"github.com/morikuni/failure"
	"os"
)

func Handle(err error) {
	code, _ := failure.CodeOf(err)
	if code == Operation {
		message, _ := failure.MessageOf(err)
		if _, err := fmt.Fprintln(os.Stderr, message); err != nil {
			panic(err)
		}
		return
	}

	fmt.Println("============ Error ============")
	fmt.Printf("Error = %v\n", err)

	fmt.Printf("Code = %v\n", code)

	msg, _ := failure.MessageOf(err)
	fmt.Printf("Message = %v\n", msg)

	cs, _ := failure.CallStackOf(err)
	fmt.Printf("CallStack = %v\n", cs)

	fmt.Printf("Cause = %v\n", failure.CauseOf(err))

	fmt.Println()
	fmt.Println("============ Detail ============")
	fmt.Printf("%+v\n", err)
}
