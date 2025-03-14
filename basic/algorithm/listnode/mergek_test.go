package listnode

import "testing"

func TestMergeKLists(t *testing.T) {
	lists := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	nodes := make([]*ListNode, len(lists))
	for i, v := range lists {
		nodes[i] = NewListNode(v)
	}
	ans := mergeKLists(nodes)
	for ans != nil {
		t.Log(ans.Val)
	}
}

func NewListNode(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head.Next
}
