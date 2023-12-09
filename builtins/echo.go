package builtins

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidArgCountEcho = errors.New("invalid argument count")
	// PlaceholderErr = errors.New("")
)

func Echo(w io.Writer, args ...string) (error) {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%w: expected at least one argument, echo is used to echo a string, need a string to echo", ErrInvalidArgCountEcho)
	default:
		ech := fmt.Sprint(args)
		// PlaceholderErr = errors.New(ech)
		return fmt.Sprintln(ech)
	}
}
