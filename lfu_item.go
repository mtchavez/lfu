package lfu

type lfuItem struct {
	data   interface{}
	parent int
}

func newlfuItem(data interface{}, parent int) *lfuItem {
	return &lfuItem{
		data:   data,
		parent: parent,
	}
}
