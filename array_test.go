package collection

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	// Test that NewCollection returns a non-nil pointer to a Collection struct
	arr := []int32{1, 2, 3}
	c := NewCollection[int32](arr)
	if c == nil {
		t.Error("NewCollection returned nil")
	}
}
