package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintCurrentDirNoArgs(t *testing.T) {
	var buf bytes.Buffer
	err := PrintCurrentDir(&buf)
	if err != NonError {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := "/home/bob/Documents/Fall-2023/CSCE-4600/A2/new/CSCE-4600-A2/Project2"
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestPrintCurrentDirLogical(t *testing.T) {
	var buf bytes.Buffer
	err := PrintCurrentDir(&buf, "-L")
	if err != NonError {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := "/home/bob/Documents/Fall-2023/CSCE-4600/A2/new/CSCE-4600-A2/Project2"
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestPrintCurrentDirPhysical(t *testing.T) {
	var buf bytes.Buffer
	err := PrintCurrentDir(&buf, "-P")
	if err != NonError {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := "/home/bob/Documents/Fall-2023/CSCE-4600/A2/new/CSCE-4600-A2/Project2"
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestPrintCurrentDirInvalidArgs(t *testing.T) {
	var buf bytes.Buffer
	err := PrintCurrentDir(&buf, "-invalidFlag")
	if err != ErrInvalidArgCountPwd {
		t.Errorf("Expected error %v, got %v", ErrInvalidArgCountPwd, err)
	}
}
