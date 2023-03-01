package solutions

// 树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap []*ListNode

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
	*h = append(*h, a.(*ListNode))
}

func (h *Heap) Pop() any {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
