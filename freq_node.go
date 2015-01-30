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
	n := &freqNode{
		value: 0,
		items: set.NewSet(),
	}
	n.prev = n
	n.next = n
	return n
}
