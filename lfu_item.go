package lfu

type lfuItem struct {
	data   interface{}
	parent *freqNode
}

func newlfuItem(data interface{}, parent *freqNode) *lfuItem {
	return &lfuItem{
		data:   data,
		parent: parent,
	}
}
