package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"testing"
	"fmt"
)


var (
	/*ErrInvalidArgCountCat = errors.New("invalid argument count")
	EmptyFileError = errors.New("invalid file")*/
)

func TestEcho(t *testing.T) {
  var out bytes.Buffer
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
		wantFile     string
		wantErr      error
		wantOut      string
	}{
		{
			name: "spaced out args",
			args: args{
				args: []string{"abc ", "def "},
			},
			wantOut: fmt.Sprintln("abc def "),
		},
		{
			name: "no args should throw invalid args error",
      			args: args{
				args: []string{},
			},
			wantErr: builtins.ErrInvalidArgCountEcho,
		},
		{
			name: "valid arguments should print text",
			args: args{
				args: []string{"\"HelloWorld\""},
			},
			wantOut: fmt.Sprintln("HelloWorld"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			/*if tt.unsetHomedir {
				oldVal := builtins.HomeDir
				t.Cleanup(func() {
					builtins.HomeDir = oldVal
				})
				builtins.HomeDir = ""
			}*/

			// testing
			if err := builtins.Echo(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("Echo() unexpected error: %v", err)
			}

			// "happy" path
      /*
			file, err := os.Open(filNam)
			if err != nil {
				t.Fatalf("Could not open temp file")
			}
			file.Close()*/
			/*d1, err := os.Stat(wd)
			if err != nil {
				t.Fatalf("Could not stat dir: %v", wd)
			}
			d2, err := os.Stat(tt.wantDir)
			if err != nil {
				t.Fatalf("Could not stat dir: %v", tt.wantDir)
			}
			if !os.SameFile(d1, d2) {
				t.Errorf("Working Directory = %v, wantDir %v", wd, tt.wantDir)
			}*/
		})
	}
}
