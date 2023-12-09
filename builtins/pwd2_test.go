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

func TestPrintCurrentDir(t *testing.T) {
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
			name: "args passed should not throw error Logical",
			args: args{
				args: []string{"-L"},
			},
			wantErr: nil,
		},
    		{
			name: "args passed should not throw error Physical",
			args: args{
				args: []string{"-P"},
			},
			wantErr: nil,
		},
		{
			name: "no args should print working directory",
        		args: args{
				args: []string{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// testing
			if err := builtins.PrintCurrentDir(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("PrintCurrentDir() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("PrintCurrentDir() unexpected error: %v", err)
			}
		})
	}
}
