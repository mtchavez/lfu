package lfu

import (
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	cache := NewLFU()
	for i := 0; i < b.N; i++ {
		cache.Insert(i, []int{i})
	}
}

func BenchmarkParallelInsert(b *testing.B) {
	cache := NewLFU()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			cache.Insert(i, []byte("user:data:cached"))
		}
	})
}

func BenchmarkGet_EmptyCache(b *testing.B) {
	cache := NewLFU()
	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}

func BenchmarkGet_AllMisses(b *testing.B) {
	cache := NewLFU()
	total := 1000000

	// Insert 1 million items
	for i := 0; i < total; i++ {
		cache.Insert(i, []int{i})
	}

	b.ResetTimer()
	// Attempt to get items that are not in cache
	for i := 0; i < b.N; i++ {
		cache.Get(i + total)
	}
}

func BenchmarkGet_AllHits(b *testing.B) {
	cache := NewLFU()
	total := 1000000

	// Insert 1 million items
	for i := 0; i < total; i++ {
		cache.Insert(i, []int{i})
	}

	b.ResetTimer()
	// Attempt to get items that are not in cache
	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}

func BenchmarkParallelGet(b *testing.B) {
	cache := NewLFU()
	total := 1000000

	// Insert 1 million items
	for i := 0; i < total; i++ {
		cache.Insert(i, []int{i})
	}

	b.ResetTimer()

	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			cache.Insert(i, []byte("user:data:cached"))
		}
	})
}
