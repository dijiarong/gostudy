package listnode

func mergeKLists(lists []*ListNode) *ListNode {
	// 初始化一个最小堆
	h := NewHeap(len(lists))
	for _, v := range lists {
		h.Push(v)
	}
	pre := &ListNode{}
	cur := pre
	for h.Peek() != nil {
		tmp := h.Pop()
		if tmp == nil {
			continue
		}
		cur.Next = tmp
		cur = cur.Next
		if tmp.Next != nil {
			h.Push(tmp.Next)
		}
	}
	return pre.Next
}

type heap struct {
	arr  []*ListNode
	size int
}

func NewHeap(cap int) *heap {
	return &heap{
		arr:  make([]*ListNode, cap),
		size: 0,
	}
}

func (h *heap) parent(index int) int {
	return (index - 1) / 2
}

func (h *heap) left(index int) int {
	return 2*index + 1
}

func (h *heap) right(index int) int {
	return 2*index + 2
}

func (h *heap) Size() int {
	return h.size
}

func (h *heap) Peek() *ListNode {
	if h.size == 0 {
		return nil
	}
	return h.arr[0]
}

func (h *heap) Push(node *ListNode) {
	// 放到最后一位
	h.arr[h.size] = node
	h.up(h.size)
	h.size++
}

func (h *heap) Pop() *ListNode {
	if h.size == 0 {
		return nil
	}
	res := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.down(0)
	return res
}

func (h *heap) up(index int) {
	for index > 0 && h.arr[index].Val < h.arr[h.parent(index)].Val {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *heap) down(index int) {
	for h.left(index) < h.size || h.right(index) < h.size {
		min := index
		if h.left(index) < h.size && h.arr[min].Val > h.arr[h.left(index)].Val {
			min = h.left(index)
		}
		if h.right(index) < h.size && h.arr[min].Val > h.arr[h.right(index)].Val {
			min = h.right(index)
		}
		if min == index {
			break
		}
		h.swap(index, min)
		index = min
	}
}

func (h *heap) swap(a, b int) {
	h.arr[a], h.arr[b] = h.arr[b], h.arr[a]
}
