package repository

import (
	"testing"
)

func TestConnection(t *testing.T) {
	conn := Connection()
	if conn == "" {
		t.Error("Expected connection string, got empty string")
	}
	if conn[:10] != "postgres://" {
		t.Error("Expected postgres connection string")
	}
}
