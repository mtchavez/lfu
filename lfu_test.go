package lfu

import (
	"reflect"
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

func TestGet(t *testing.T) {
	cache := NewLFU()
	expectedData := []byte("user-signout:34:987654321")
	cache.Insert(2, expectedData)
	data, e := cache.Get(2)
	if e != nil {
		t.Fatalf("Got unexpected error on Get: %+v", e)
	}
	if !reflect.DeepEqual(data, expectedData) {
		t.Fatalf("Expected %+v data to be returned but got %+v", expectedData, data)
	}
}

func TestGet_notFound(t *testing.T) {
	cache := NewLFU()
	data, e := cache.Get(2)
	if e == nil {
		t.Fatalf("Expected an error but got nil")
	}
	if data != nil {
		t.Fatalf("Expected no data to be found but got %+v", data)
	}
}
