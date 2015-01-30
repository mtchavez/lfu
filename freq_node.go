package lfu

import (
	set "github.com/deckarep/golang-set"
)

type freqNode struct {
	value int
	items set.Set
	prev  *freqNode
	next  *freqNode
}

func newFreqNode() *freqNode {
	return &freqNode{
		value: 0,
		items: set.NewSet(),
	}
}
