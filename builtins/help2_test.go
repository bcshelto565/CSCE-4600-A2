package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"testing"
	"fmt"
)

var list string

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
		list	string
	}{
		{
			name: "3 args passed should throw error",
			args: args{
				args: []string{"abc ", "def ", "ghi "},
			},
			wantErr: builtins.ErrInvalidArgCountHelp,
		},
		{
			name: "2 args passed should throw error",
			args: args{
				args: []string{"abc ", "def "},
			},
			wantErr: builtins.ErrInvalidArgCountHelp,
		},
		{
			name: "1 arg passed should throw error",
			args: args{
				args: []string{"abc "},
			},
			wantErr: builtins.ErrInvalidArgCountHelp,
		},
		{
			name: "no args should print helpList with no errors",
        		args: args{
				args: []string{},
			},
			wantErr: nil,
		},
		{
			name: "no args should print helpList",
        		args: args{
				args: []string{},
			},
			list: ("alias\ncat\ncd\necho\nenv\nhelp\npwd\n"),
			wantOut: fmt.Sprintln(list),
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
