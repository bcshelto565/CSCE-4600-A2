package builtins_test

import (
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"os"
	"testing"
)

func TestCommandAlias(t *testing.T) {
	tmp := t.TempDir()

	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
		wantList     bool

		wantErr      error
	}{
		{
			name: "error too many args",
			args: args{
				args: []string{"abc", "def", "ghi", "jkl"},
			},
			wantErr: builtins.ErrInvalidArgCountAlias,
		},
		{
			name:    "no args should throw error",
			wantErr: builtins.ErrInvalidArgCountAlias,		
		},
		{
			name: "one arg should print aliases if arg is -p",
			args: args{
				args: []string{"-p"},
			},
			wantList: true
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			if tt.unsetHomedir {
				oldVal := builtins.HomeDir
				t.Cleanup(func() {
					builtins.HomeDir = oldVal
				})
				builtins.HomeDir = ""
			}

			// testing
			if err := builtins.CommandAlias(tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ChangeDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("ChangeDirectory() unexpected error: %v", err)
			}

			// "happy" path
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("Could not get working dir")
			}
			d1, err := os.Stat(wd)
			if err != nil {
				t.Fatalf("Could not stat dir: %v", wd)
			}
			d2, err := os.Stat(tt.wantDir)
			if err != nil {
				t.Fatalf("Could not stat dir: %v", tt.wantDir)
			}
			if !os.SameFile(d1, d2) {
				t.Errorf("Working Directory = %v, wantDir %v", wd, tt.wantDir)
			}
		})
	}
}