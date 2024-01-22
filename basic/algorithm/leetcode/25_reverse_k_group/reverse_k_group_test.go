package reversekgroup

import "testing"

func TestReverseKGroup(t *testing.T) {
	root := initNode([]int{1, 2, 3, 4, 5})
	reverseKGroup(root, 2)
}

func initNode(slice []int) *ListNode {
	pre := &ListNode{}
	pro := pre
	for _, v := range slice {
		cur := &ListNode{v, nil}
		pre.Next = cur
		pre = cur
	}
	return pro.Next
}
