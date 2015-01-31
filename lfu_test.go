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

func TestGetLFUItem(t *testing.T) {
	cache := NewLFU()
	cache.Insert(1, "guest:987654321")
	cache.Insert(2, "user:42:123456789")
	cache.Get(1)
	cache.Get(1)
	cache.Get(2)
	value, data := cache.GetLFUItem()
	if value != 2 {
		t.Fatalf("Expected 2 to be LFU but got %+v", value)
	}
	if !reflect.DeepEqual(data, "user:42:123456789") {
		t.Fatalf("Expected data for 2 to be returned but got %+v", data)
	}
}

func TestGetLFUItem_emptyCache(t *testing.T) {
	cache := NewLFU()
	value, data := cache.GetLFUItem()
	if value != -1 {
		t.Fatalf("Expected -1 to be LFU for empty cache but got %+v", value)
	}
	if data != nil {
		t.Fatalf("Expected data to be nil but got %+v", data)
	}
}
