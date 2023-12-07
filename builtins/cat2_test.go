package builtins_test

import (
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"os"
	"testing"
	"io/ioutil"
	// "./cat.go"
	// "builtins"
)


var (
	/*ErrInvalidArgCountCat = errors.New("invalid argument count")
	EmptyFileError = errors.New("invalid file")*/
)

func TestPrintFileContents(t *testing.T) {
	// tmp := t.TempDir()
	tmpfile, temperr := ioutil.TempFile("", "example")
	if temperr != nil {
		t.Fatalf("tempfile not opened, error = %v", temperr)
	}
	tempstring := "example"
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
		// unsetHomedir bool
		wantFile     string
		wantErr      error
	}{
		{
			name: "error too many args",
			args: args{
				args: []string{"abc", "def"},
			},
			wantErr: builtins.ErrInvalidArgCountCat,
		},
		{
			name:    "no args should throw invalid args error",
      args: args{
				args: []string{" "},
			},
			wantErr: builtins.ErrInvalidArgCountCat,
		},
		{
			name:         "no file contents should throw empty file error",
			wantErr:      builtins.EmptyFileError,
		},
		{
			name: "one arg should print file contents",
			args: args{
				args: []string{tmpfile.Name},
			},
			wantFile: tempstring,
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
			if err := builtins.PrintFileContents(tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("PrintFileContents() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("PrintFileContents() unexpected error: %v", err)
			}

			// "happy" path
			file, err := os.Open(tmpfile)
			if err != nil {
				t.Fatalf("Could not open temp file")
			}
			file.Close()
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
