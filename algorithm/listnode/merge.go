package listnode

type mergeListNode struct{}

var MergeListNode mergeListNode

// 合并两个有序的链表
func (mergeListNode) Merge(list1 *ListNode, list2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}
	return pre.Next
}
