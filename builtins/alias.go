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

type ComAlias struct {			// struct for new custom alias entries
	name string
	value string
}

aliasSlic := []ComAlias{		// slice for holding all the custom alias entries
}

func printComs(aliasSlic []ComAlias) {		// simple for loop function to print entries
	for _, ComAlias := range aliasSlic {
		fmt.Fprintln("Alias name: %s, Alias value: %s", ComAlias.name, ComAlias.value)
	}
}

func CommandAlias(args ...string) error {
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%w: expected at least one argument, one for print or write of a name, and one for value of alias", ErrInvalidArgCount)
	case 1:
		if args[0] == "-p" {
			printComs(ComAlias struct)
		}
	case 2:
		return fmt.Errorf("%w: missing arguments, alias needs to be used in the following syntax: alias update = \"sudo apt-get update\" ensure the \"\" quotes are used to define the command", ErrInvalidArgCount)
	case 3:
		aliasSlic = append(aliasSlic, ComAlias{name: arg[0], value: arg[2]})
		fmt.Println("new alias is: ", arg[0], " = ", arg[2])
	default:
		return fmt.Errorf("%w: Expected 1 or 3 arguments, 1 argument of \"-p\" to print alias list, and 3 arguments for an alias entry.", ErrInvalidArgCount)
	}
}
