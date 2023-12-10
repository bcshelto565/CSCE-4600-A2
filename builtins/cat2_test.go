package builtins_test

import (
	"bytes"
	"errors"
	"github.com/bcshelto565/CSCE-4600-A2/builtins"
	"os"
	"testing"
)

func TestPrintFileContents(t *testing.T) {
	var out bytes.Buffer
	tmpfile, temperr := os.CreateTemp("", "example")
	if temperr != nil {
		t.Fatalf("tempfile not opened, error = %v", temperr)
	}
	tempstring := "example"
	filNam := tmpfile.Name()
	tmpfile2, temperr2 := os.CreateTemp("", "example2")
	if temperr2 != nil {
		t.Fatalf("tempfile not opened, error = %v", temperr)
	}
	_, writeerr := tmpfile.WriteString("example")
	if writeerr != nil {
		t.Fatalf("tempfile not written to, error = %v", writeerr)
	}
	filNam2 := tmpfile2.Name()
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
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
		})
	}
}
