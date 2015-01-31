package lfu

import (
	"errors"
	"fmt"
)

type LFUCache struct {
	frequencies map[int]*lfuItem
	head        *freqNode
}

func NewLFU() *LFUCache {
	return &LFUCache{
		frequencies: make(map[int]*lfuItem),
		head:        newFreqNode(),
	}
}

func (c *LFUCache) Insert(key int, value interface{}) (bool, error) {
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

func (c *LFUCache) Get(key int) (interface{}, error) {
	item, existing := c.frequencies[key]
	if !existing {
		return nil, errors.New(fmt.Sprintf("Key: %+v not found in cache", key))
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
