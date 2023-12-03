package builtins

import (
	"bytes"
	"testing"
)

var str string

func stringUser(str string) {
	str += str
	return
}

func TestCommandAliasPrintAliases(t *testing.T) {
	var buf bytes.Buffer
	str, err := CommandAlias(&buf, "-p")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	stringUser(str)
	expectedOutput := ""
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestCommandAliasAddAlias(t *testing.T) {
	var buf bytes.Buffer
	str, err := CommandAlias(&buf, "newAlias = whoami")
	err = nil
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	stringUser(str)
	var expectedOutput string
	expectedOutput = "Alias name: newAlias, Alias value: whoami"
	got := buf.String()
	if got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestCommandAliasInvalidArgs(t *testing.T) {
	var buf bytes.Buffer
	str, err := CommandAlias(&buf)
	if err != ErrInvalidArgCountAlias {
		t.Errorf("Expected error %v, got %v", ErrInvalidArgCountAlias, err)
	}
	stringUser(str)
}
