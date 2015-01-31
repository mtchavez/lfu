package lfu

import (
	"errors"
	"fmt"
)

// Cache used to call core cache methods off of
type Cache struct {
	frequencies map[interface{}]*lfuItem
	head        *freqNode
}

// NewLFU initializes a new LFU cache and returns it
func NewLFU() *Cache {
	return &Cache{
		frequencies: make(map[interface{}]*lfuItem),
		head:        newFreqNode(),
	}
}

// Insert takes a key and a value to insert into the cache
// Returns a boolean for successful insert and an error if failed
func (c *Cache) Insert(key interface{}, value interface{}) (bool, error) {
	_, found := c.frequencies[key]
	if found {
		return false, errors.New("Key already exists in cache")
	}
	freq := c.head.next
	if freq.value != 1 {
		freq = getNewNode(1, c.head, freq)
	}
	freq.items.Add(key)
	c.frequencies[key] = newlfuItem(value, freq)
	return true, nil
}

// Get takes a key for an item in the cache to look up
// Returns the data associated with that key and an
func (c *Cache) Get(key interface{}) (interface{}, error) {
	item, existing := c.frequencies[key]
	if !existing {
		return nil, fmt.Errorf("Key: %+v not found in cache", key)
	}
	freq := item.parent
	nextFreq := freq.next

	if nextFreq == c.head || nextFreq.value != freq.value+1 {
		nextFreq = getNewNode(freq.value+1, freq, nextFreq)
	}
	nextFreq.items.Add(key)
	item.parent = nextFreq

	freq.items.Remove(key)
	if freq.items.Cardinality() == 0 {
		freq.remove()
	}
	return item.data, nil
}

// GetLFUItem returns the key and data of the least
// frequently updated item in the cache or -1 and nil for not found
func (c *Cache) GetLFUItem() (value interface{}, data interface{}) {
	if len(c.frequencies) == 0 {
		return -1, nil
	}
	// TODO: Try to avoid ToSlice on all items
	item := c.frequencies[c.head.next.items.ToSlice()[0]]
	return item.parent.value, item.data
}
