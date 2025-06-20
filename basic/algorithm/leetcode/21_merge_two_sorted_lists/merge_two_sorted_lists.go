package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node1, node2 := list1, list2
	pre := &ListNode{}
	cur := pre
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else {
			cur.Next = node2
			node2 = node2.Next
		}
		cur = cur.Next
	}
	if node1 != nil {
		cur.Next = node1
	}
	if node2 != nil {
		cur.Next = node2
	}
	return pre.Next
}
