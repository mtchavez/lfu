package lfu

import (
	"testing"
)

func Test_newFreqNode(t *testing.T) {
	fn := newFreqNode()
	if fn.value != 0 {
		t.Fatalf("Expected initial value to be 0 but got %+v", fn.value)
	}
	size := fn.items.Cardinality()
	if size != 0 {
		t.Fatalf("Expected items to be empty but got %+v", size)
	}
	if fn.prev != fn || fn.next != fn {
		t.Fatalf("Next and previous nodes should be self got prev: %+v and next: %+v", fn.prev, fn.next)
	}
}

func Test_freqNode_remove(t *testing.T) {
	node1 := newFreqNode()
	node2 := newFreqNode()
	node3 := newFreqNode()
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2
	node2.remove()
	if node1.next != node3 && node3.prev != node1 {
		t.Fatalf("Expected node 2 to be removed and node1 and 3 to be linked")
	}
}
