package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"testing"
	// "fmt"
)


var (
	/*ErrInvalidArgCountCat = errors.New("invalid argument count")
	EmptyFileError = errors.New("invalid file")*/
)

func TestHelp(t *testing.T) {
  var out bytes.Buffer
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
    		missinB      bool
		wantFile     string
		wantErr      error
		wantOut      string
	}{
		{
			name: "args passed should throw error",
			args: args{
				args: []string{"abc ", "def "},
			},
			wantErr: builtins.ErrInvalidArgCountHelp,
		},
		{
			name: "no args should print helpList",
        		args: args{
				args: []string{},
			},
			wantErr: nil,
		},
		/*{
			name: "builtins not found should throw missing dir error",
			missinB: true,
        		args: args{
				  args: []string{},
			},
			wantErr: builtins.ErrMissingBuiltins,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*if tt.missinB {
				oldVal := builtins.HomeDir
				t.Cleanup(func() {
					builtins.HomeDir = oldVal
				})
				builtins.HomeDir = ""
			}*/
			// testing
			if err := builtins.Help(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Help() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("Help() unexpected error: %v", err)
			}
		})
	}
}
