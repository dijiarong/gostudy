package reversekgroup

func rr(head *ListNode, k int) *ListNode {
	pro := &ListNode{
		Next: head,
	}
	pre := pro
	for head != nil {
		end := finda(head, k)
		if end == nil {
			break
		}
		nextHead := end.Next

		r1(head, nextHead)

		pre.Next = end
		head.Next = nextHead
		pre = head
		head = nextHead
	}
	return pro.Next
}

func r1(head, end *ListNode) {
	node := head
	pre := &ListNode{}
	for node != end {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
}

func finda(node *ListNode, k int) *ListNode {
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
