package lfu

import (
	"fmt"
	"reflect"
)

func ExampleCache_Insert() {
	cache := NewLFU()
	inserted, _ := cache.Insert(1, []byte("user:1:aeb31da42fae09afqld"))
	fmt.Println("Inserted:", inserted)

	// Insert existing key
	_, err := cache.Insert(1, "boom")
	fmt.Println("Insert error:", err)
	// Output:
	// Inserted: true
	// Insert error: Key already exists in cache
}

func ExampleCache_Get() {
	cache := NewLFU()
	cache.Insert(1, []byte("user:1:aeb31da42fae09afqld"))

	// Get key 1
	data, _ := cache.Get(1)
	fmt.Println("Cached data for key 1:", string(reflect.ValueOf(data).Bytes()))

	// Get a not cached key
	_, err := cache.Get(42)
	fmt.Println("Failed get:", err)
	// Output:
	// Cached data for key 1: user:1:aeb31da42fae09afqld
	// Failed get: Key: 42 not found in cache
}
