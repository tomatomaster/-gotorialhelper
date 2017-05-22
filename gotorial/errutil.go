package gotorial

import "os"
import "fmt"

//OSExitIfError exit app if err != nil
func OSExitIfError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

//OSExitWithCauseIfError print cause and exit app if err != nil
func OSExitWithCauseIfError(cause string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v %v", cause, err)
		os.Exit(1)
	}
}
