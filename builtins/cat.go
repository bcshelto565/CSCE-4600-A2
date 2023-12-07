package builtins

import (
	"errors"
	"fmt"
	"os"
	"io"
)

var (
	ErrInvalidArgCountCat = errors.New("invalid argument count")
	EmptyFileError = errors.New("invalid file")
)

func PrintFileContents(w io.Writer, args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%v", ErrInvalidArgCountCat)
	case 1:
		fil, err := os.ReadFile(args[0])
		if err != (nil) {
			return fmt.Errorf("%w: File is empty and cannot be read", EmptyFileError)
		}
		file_stat, staterr := os.Stat(args[0])
		if file_stat.Size == 0 {
			return fmt.Errorf("%w: File is empty and cannot be read", EmptyFileError)
		}
		fmt.Print(string(fil))
		return (nil)
	default:
		return fmt.Errorf("%w: one argument for file to be read, proper usage is \"cat file.txt\" where \"file.txt\" is the name of the file to be read", ErrInvalidArgCountCat)
	}
}
