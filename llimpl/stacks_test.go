package ds

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	_ = NewStack()
	if head != nil {
		t.Error("Failed to initialize stack")
	}
}
