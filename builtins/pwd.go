package builtins

import (
	"errors"
	"fmt"
	"os"
	"io"
)

var (
	ErrInvalidArgCountPwd = errors.New("invalid argument count")
)

func PrintCurrentDir(w io.Writer, args ...string) error {
	switch len(args) {
	case 0: // print current working directory
		fmt.Println(os.Getwd())
		return fmt.Errorf("%w", NonError)
	case 1:
		if (args[0] == "-L") {
			fmt.Println(os.Getwd())
			return fmt.Errorf("%w", NonError)
		} else if (args[0] == "-P") {
			fmt.Println(os.Getwd())
			return fmt.Errorf("%w", NonError)
		} else {
			return fmt.Errorf("%w: expected zero or one arguments -L for logical working directory address and -P for physical working directory address", ErrInvalidArgCountPwd)
		}
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCountPwd)
	}
}