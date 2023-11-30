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

func PrintCurrentDir(args ...string) error {
	switch len(args) {
	case 0: // print current working directory
		fmt.Println(os.Getwd())
	case 1:
		if (arg[0] == "-L") {
			fmt.Println(os.Getwd())
		}
		else if (arg[0] == "-P") {
			fmt.Println(os.Getwd())
		}
		else {
			return fmt.Errorf("%w: expected zero or one arguments -L for logical working directory address and -P for physical working directory address", ErrInvalidArgCount)
		}
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)
	}
}