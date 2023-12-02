package builtins

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestPrintFileContentsValidFile(t *testing.T) {
	var buf bytes.Buffer
	content := []byte("Hello, this is a test file.")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	err = PrintFileContents(&buf, tmpfile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := "Hello, this is a test file."
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestPrintFileContentsEmptyFile(t *testing.T) {
	var buf bytes.Buffer
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	err = PrintFileContents(&buf, tmpfile.Name())
	if err != EmptyFileError {
		t.Errorf("Expected error %v, got %v", EmptyFileError, err)
	}
}

func TestPrintFileContentsInvalidArgs(t *testing.T) {
	var buf bytes.Buffer
	err := PrintFileContents(&buf)
	if err != ErrInvalidArgCountCat {
		t.Errorf("Expected error %v, got %v", ErrInvalidArgCountCat, err)
	}
}

