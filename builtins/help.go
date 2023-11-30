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

HelpList = []string {
}

func Help(args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		dir, err := os.Open(".")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, file := range files {
			if strings.contains(file, "test"){
				continue
			}
			else {
				fileNam := file[:len(file)-3]
				HelpList = append(HelpList, fileNam)
			}
		}
		for i, s := range HelpList {
			fmt.Println(s)
		}
	}

}