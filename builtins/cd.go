package builtins

import (
	"errors"
	"fmt"
	"os"
	"github.com/jh125486/CSCE4600/Project2/builtins"
)

var (
	ErrInvalidArgCountCd = errors.New("invalid argument count")
	HomeDir, _         = os.UserHomeDir()
)

func ChangeDirectory(args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available
		if HomeDir == "" {
			return fmt.Errorf("%w: no homedir found, expected one argument (directory)", builtins.ErrInvalidArgCount)
		}
		return os.Chdir(HomeDir)
	case 1:
		return os.Chdir(args[0])
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", builtins.ErrInvalidArgCount)
	}
}
