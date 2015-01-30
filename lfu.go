package lfu

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
