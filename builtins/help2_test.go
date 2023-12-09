package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"testing"
	// "fmt"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
