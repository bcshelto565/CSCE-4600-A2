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

func PrintFileContents(args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%w: expected at least one argument, cat is used to print a file's contents, need a file to print", ErrInvalidArgCount)
	case 1:
		fil, err := os.ReadFile(arg[0])
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(fil)
	default:
		return fmt.Errorf("%w: one argument for file to be read, proper usage is \"cat file.txt\" where \"file.txt\" is the name of the file to be read", ErrInvalidArgCount)
	}
}