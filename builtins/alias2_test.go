package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"testing"
	"fmt"
)

func TestCommandAlias(t *testing.T) {
  var out bytes.Buffer
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
    		listExist    bool
		wantFile     string
		wantErr      error
		wantOut string
	}{
		{
			name: "error too many args",
			args: args{
				args: []string{"abc ", "= ", "echo ", "def "},
			},
			wantErr: builtins.ErrInvalidArgCountAlias,
		},
		{
			name:    "no args should throw invalid args error",
      			args: args{
				args: []string{},
			},
			wantErr: builtins.ErrInvalidArgCountAlias,
		},
		{
			name:         "no existing aliases should print nothing",
			args: args{
				args: []string{"-p"},
			},
			wantOut: fmt.Sprintln(""),
		},
		{
			name: "valid arguments should add to list",
			args: args{
				args: []string{"ech ", "= ", "\"echo\""},
			},
			wantOut: fmt.Sprintln("new alias is: ech = echo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			// testing
			if err := builtins.CommandAlias(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("CommandAlias() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("CommandAlias() unexpected error: %v", err)
			}
		})
	}
}
