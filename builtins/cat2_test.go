package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"os"
	"testing"
	"io"
)


var (
	/*ErrInvalidArgCountCat = errors.New("invalid argument count")
	EmptyFileError = errors.New("invalid file")*/
)

func TestPrintFileContents(t *testing.T) {
	// tmp := t.TempDir()
	var out bytes.Buffer
	tmpfile, temperr := ioutil.TempFile("", "example")
	if temperr != nil {
		t.Fatalf("tempfile not opened, error = %v", temperr)
	}
	tempstring := "example"
	filNam := tmpfile.Name()
	tmpfile2, temperr2 := ioutil.TempFile("", "example2")
	if temperr2 != nil {
		t.Fatalf("tempfile not opened, error = %v", temperr)
	}
	tmpfile.WriteString("example")
	filNam2 := tmpfile2.Name()
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
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name:    "no args should throw invalid args error",
      			args: args{
				args: []string{},
			},
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name:         "no file contents should throw empty file error",
			args: args{
				args: []string{filNam2},
			},
			wantErr:      builtins.EmptyFileError,
		},
		{
			name: "one arg should print file contents",
			args: args{
				args: []string{filNam},
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
			if err := builtins.PrintFileContents(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("PrintFileContents() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("PrintFileContents() unexpected error: %v", err)
			}

			// "happy" path
			file, err := os.Open(filNam)
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
