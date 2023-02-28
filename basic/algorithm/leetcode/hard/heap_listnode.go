package hard

import (
	"gostudy/basic/algorithm/listnode"
)

type Heap []*listnode.ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	tmp := h[j]
	h[j] = h[i]
	h[i] = tmp
}

func (h *Heap) Push(a any) {
	*h = append(*h, a.(*listnode.ListNode))
}

func (h *Heap) Pop() any {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
