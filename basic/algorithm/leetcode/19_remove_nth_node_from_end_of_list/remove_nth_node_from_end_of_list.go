package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNodes(nums []int) *ListNode {
	head := &ListNode{
		Val: nums[0],
	}
	cur := head
	for i := 1; i < len(nums); i++ {
		tmp := &ListNode{
			Val: nums[i],
		}
		cur.Next = tmp
		cur = tmp
	}
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	tmp := findNthFromEnd(pre, n+1)
	if tmp.Next != nil {
		tmp.Next = tmp.Next.Next
	}
	return pre.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head.Next
	for fast != nil {
		n--
		if n == 0 {
			break
		}
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
