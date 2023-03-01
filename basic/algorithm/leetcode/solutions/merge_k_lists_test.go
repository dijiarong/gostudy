package solutions

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{}
	node6 := &ListNode{Val: 5, Next: nil}
	node5 := &ListNode{Val: 3, Next: node6}
	node4 := &ListNode{Val: 1, Next: node5}
	node3 := &ListNode{Val: 6, Next: nil}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 2, Next: node2}
	lists = append(lists, node1)
	lists = append(lists, node4)
	MergeKLists(lists)
}
