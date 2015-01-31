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

func TestInsert(t *testing.T) {
	cache := NewLFU()
	success, e := cache.Insert(2, []byte("user-login:34:123456789"))
	if !success || e != nil {
		t.Fatalf("Expected to be inserted into cache but got %+v - %+v", success, e)
	}
}

func TestInsert_duplicate(t *testing.T) {
	cache := NewLFU()
	cache.Insert(2, []byte("user-signout:34:987654321"))
	success, e := cache.Insert(2, []byte("user-login:34:123456789"))
	if success || e == nil {
		t.Fatalf("Expected error trying to insert duplicate key")
	}
}
