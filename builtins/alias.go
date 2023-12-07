package builtins

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	ErrInvalidArgCountAlias = errors.New("invalid argument count: Expected 1 or 3 arguments, 1 argument of \"-p\" to print alias list, and 3 arguments for an alias entry.")
	NonError = errors.New("")
)

type ComAlias struct {			// struct for new custom alias entries
	name string
	value string
}

var aliasSlic []ComAlias		// slice for holding all the custom alias entries

func printComs(aliasSlic []ComAlias) string {		// simple for loop function to print entries
	var output string
	for _, ComAlias := range aliasSlic {
		nam := ComAlias.name
		val := ComAlias.value
		output += "Alias name: " + nam + ", Alias value: " + val + "\n"
	}
	return output
}

func CommandAlias(w io.Writer, args ...string) (error) {
	// var output string
	switch len(args) {		// switch case for determining how to run the command
	case 0:
		return fmt.Errorf("%w", ErrInvalidArgCountAlias)
	case 1:
		if args[0] == "-p" {
			// output := string(printComs(aliasSlic))
			output := make([]string, 0)
			output = append(output, string(printComs(aliasSlic)))
			_, err := fmt.Fprintln(w, strings.Join(output, "\n"))
			return err
		} else {
			return fmt.Errorf("%w", ErrInvalidArgCountAlias)
			// _, err := fmt.Fprintln(w, strings.Join(toShow, "\n"))
		}
	case 2:
		return fmt.Errorf("%w: missing arguments, alias needs to be used in the following syntax: alias update = \"sudo apt-get update\" ensure the \"\" quotes are used to define the command", ErrInvalidArgCountAlias)
	case 3:
		aliasSlic = append(aliasSlic, (ComAlias{name: args[0], value: args[2]}))
		// fmt.Println("new alias is: ", args[0], " = ", args[2])
		arg1 := args[0]
		arg2 := args[2]
		var output string
		output = "new alias is: " + arg1 + " = " + arg2
		/*output := make([]string, 0)
		output = append(output, "new alias is: ", arg1, " = ", arg2)*/
		// output += "new alias is: " + arg1 + " = " + arg2
		_, err := fmt.Fprintln(w, strings.Join(output, "\n"))
		return err
	default:
		return fmt.Errorf("%w: Expected 1 or 3 arguments, 1 argument of \"-p\" to print alias list, and 3 arguments for an alias entry.", ErrInvalidArgCountAlias)
	}
}
