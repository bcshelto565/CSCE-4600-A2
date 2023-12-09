package builtins

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"io"
)

var (
	ErrInvalidArgCountHelp = errors.New("invalid argument count: expected no arguments for help. help prints out available builtin commands")
	ErrMissingBuiltins = errors.New("cannot find builtins directory, cannot list builtins")
)

var HelpList []string

func Help(w io.Writer, args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		dir, err := os.Open("./builtins")
		if err != nil {
			fmt.Println("Error:", err)
			return fmt.Errorf("%w: Directory failed to open", ErrMissingBuiltins)
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			fmt.Println("Error:", err)
			return fmt.Errorf("%w: Directory failed to open", EmptyFileError)
		}
		for _, file := range files {
			if strings.Contains((file.Name()), "test"){
				continue
			} else {
				fileNam := ((file.Name())[:len(file.Name())-3])
				HelpList = append(HelpList, fileNam)
			}
		}
		helpList := make([]string, 0)
		helpList = append(helpList, HelpList...)
		_, err2 := fmt.Fprintln(w, strings.Join(helpList, "\n"))
		return err2
	default:
		return fmt.Errorf("%w", ErrInvalidArgCountHelp)
	}
}
