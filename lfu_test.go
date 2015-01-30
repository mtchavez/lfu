package lfu

import (
	"testing"
)

func TestNewLFU(t *testing.T) {
	cache := NewLFU()
	if len(cache.frequencies) != 0 {
		t.Fatalf("Expected no frequencies on new cache got %+v", len(cache.frequencies))
	}
}
