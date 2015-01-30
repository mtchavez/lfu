package lfu

import (
	"testing"
)

func Test_newlfuItem(t *testing.T) {
	data := "My Data"
	parent := 42
	item := newlfuItem(data, parent)
	if item.data != data {
		t.Fatalf("Expected data to be %+v but got %+v", data, item.data)
	}
	if item.parent != parent {
		t.Fatalf("Expected parent to be %+v but got %+v", parent, item.parent)
	}
}
