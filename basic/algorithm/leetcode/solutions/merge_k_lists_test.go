package solutions

import (
	"gostudy/basic/algorithm/listnode"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*listnode.ListNode{}
	node6 := &listnode.ListNode{Val: 5, Next: nil}
	node5 := &listnode.ListNode{Val: 3, Next: node6}
	node4 := &listnode.ListNode{Val: 1, Next: node5}
	node3 := &listnode.ListNode{Val: 6, Next: nil}
	node2 := &listnode.ListNode{Val: 4, Next: node3}
	node1 := &listnode.ListNode{Val: 2, Next: node2}
	lists = append(lists, node1)
	lists = append(lists, node4)
	MergeKLists(lists)
}
