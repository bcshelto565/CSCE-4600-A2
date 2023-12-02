package builtins

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEchoWithArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Echo(&buf, "Hello", "World", "123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOutput := "Hello World 123"
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestEchoNoArgs(t *testing.T) {
	var buf bytes.Buffer
	err := Echo(&buf)
	if err != ErrInvalidArgCountEcho {
		t.Errorf("Expected error %v, got %v", ErrInvalidArgCountEcho, err)
	}
}

