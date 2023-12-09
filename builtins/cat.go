package builtins

import (
	"errors"
	"fmt"
	"os"
	"io"
	"github.com/jh125486/CSCE4600/Project2/builtins"
)

var (
	EmptyFileError = errors.New("invalid file")
)

func PrintFileContents(w io.Writer, args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%v", builtins.ErrInvalidArgCount)
	case 1:
		fil, err := os.ReadFile(args[0])
		if err != (nil) {
			return fmt.Errorf("%w: File is empty and cannot be read", EmptyFileError)
		}
		file_stat, staterr := os.Stat(args[0])
		file_size := file_stat.Size()
		if staterr != (nil) {
			return fmt.Errorf("%w: File is empty and cannot be read", EmptyFileError)
		}
		if file_size == 0 {
			return fmt.Errorf("%w: File is empty and cannot be read", EmptyFileError)
		}
		fmt.Print(string(fil))
		return (nil)
	default:
		return fmt.Errorf("%w", builtins.ErrInvalidArgCount)
	}
}
