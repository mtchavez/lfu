package lfu

import "errors"

type LFUCache struct {
	frequencies map[int]interface{}
	head        *freqNode
}

func NewLFU() *LFUCache {
	return &LFUCache{
		frequencies: make(map[int]interface{}),
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
	c.frequencies[key] = value
	return true, nil
}
