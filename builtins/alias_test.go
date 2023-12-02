package builtins

import (
	"bytes"
	"testing"
)

func TestCommandAliasPrintAliases(t *testing.T) {
	var buf bytes.Buffer
	err := CommandAlias(&buf, "-p")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := ""
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestCommandAliasAddAlias(t *testing.T) {
	var buf bytes.Buffer
	err := CommandAlias(&buf, "newAlias", "=", "echo \"Hello, World!\"")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := "new alias is: newAlias = echo \"Hello, World!\""
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestCommandAliasInvalidArgs(t *testing.T) {
	var buf bytes.Buffer
	err := CommandAlias(&buf)
	if err != ErrInvalidArgCountAlias {
		t.Errorf("Expected error %v, got %v", ErrInvalidArgCountAlias, err)
	}
}
