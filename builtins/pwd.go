package builtins

import (
	"errors"
	"fmt"
	"os"
	"io"
	"strings"
)

var (
	ErrInvalidArgCountPwd = errors.New("invalid argument count")
)

func PrintCurrentDir(w io.Writer, args ...string) error {
	switch len(args) {
	case 0: // print current working directory
		pwd, _ := os.Getwd()
		output := make([]string, 0)
		output = append(output, pwd)
		_, err := fmt.Fprintln(w, strings.Join(output, "\n"))
		return err
	case 1:
		if (args[0] == "-L") {
			pwd, _ := os.Getwd()
			output := make([]string, 0)
			output = append(output, pwd)
			_, err := fmt.Fprintln(w, strings.Join(output, "\n"))
			return err
		} else if (args[0] == "-P") {
			pwd, _ := os.Getwd()
			output := make([]string, 0)
			output = append(output, pwd)
			_, err := fmt.Fprintln(w, strings.Join(output, "\n"))
			return err
		} else {
			return fmt.Errorf("%w: expected zero or one arguments -L for logical working directory address and -P for physical working directory address", ErrInvalidArgCount)
		}
	default:
		return fmt.Errorf("%w", ErrInvalidArgCount)
	}
}
