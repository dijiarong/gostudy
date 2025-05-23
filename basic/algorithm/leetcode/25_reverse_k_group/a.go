package reversekgroup

func r(head *ListNode, k int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	cur := pre
	for head != nil {
		end := find(head, k)
		if end == nil {
			break
		}
		endNext := end.Next
		re(head, endNext)

		cur.Next = end
		head.Next = endNext
		cur = head
		head = endNext
	}
	return pre.Next
}

func re(node, end *ListNode) {
	tmp := node
	var pre *ListNode
	for tmp != end {
		next := tmp.Next
		tmp.Next = pre
		pre = tmp
		tmp = next
	}
}

func find(node *ListNode, k int) *ListNode {
	if node == nil {
		return nil
	}
	for node != nil {
		k--
		if k == 0 {
			return node
		}
		node = node.Next
	}
	return nil
}
