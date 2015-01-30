package lfu

import (
	"testing"
)

func TestNewFreqNode(t *testing.T) {
	fn := newFreqNode()
	if fn.value != 0 {
		t.Fatalf("Expected initial value to be 0 but got %+v", fn.value)
	}
	size := fn.items.Cardinality()
	if size != 0 {
		t.Fatalf("Expected items to be empty but got %+v", size)
	}
}
