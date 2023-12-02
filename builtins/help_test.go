package builtins

import (
	"bytes"
	"testing"
)

func TestHelpNoArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Help(&buf)
	if err != nil {
		t.Error("Expected no error, got %w", err)
	}
	expectedOutput := "cd\ncat\npwd\nalias\nhelp\nenv\necho\n"
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output: %q, got %q", expectedOutput, got)
	}
}

func TestHelpWithArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Help(&buf, "arg1", "arg2")
	if err == nil {
		t.Error("Expected error of too many arguments, got nil")
	}
	expectedError := ErrInvalidArgCountHelp
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}
