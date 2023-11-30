package builtins

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
	HomeDir, _         = os.UserHomeDir()
)

func Echo(args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%w: expected at least one argument, echo is used to echo a string, need a string to echo", ErrInvalidArgCount)
	default:
		ech := fmt.Sprint(args)
		fmt.Fprintln(ech)
	}
}