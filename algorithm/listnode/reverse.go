package listnode

type reverseListNode struct{}

var ReverseListNode reverseListNode

// 反转链表（普通版）
func (reverseListNode) Reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// 反转链表 （递归）
func (reverseListNode) ReverseRec(head *ListNode) *ListNode {
	return recursion(nil, head)
}

func recursion(pre *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return recursion(cur, next)
}
