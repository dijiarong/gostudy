package listnode

type palindromeListNode struct{}

var PalindromeListNode palindromeListNode

// 记录数值,空间O(n) 时间O(n)
func (palindromeListNode) IsPalindrome1(head *ListNode) bool {
	valList := []int{}
	cur := head
	for cur != nil {
		valList = append(valList, cur.Val)
		cur = cur.Next
	}
	for i := len(valList) - 1; head != nil; i-- {
		if head.Val != valList[i] {
			return false
		}
		head = head.Next
	}
	return true
}

// 反转后半部分
func (palindromeListNode) IsPalindrome2(head *ListNode) bool {
	left := head
	right := head
	for right.Next != nil && right.Next.Next != nil {
		right = right.Next.Next
		left = left.Next
	}
	if left.Next == nil {
		return true
	}
	// 反转
	rightHalf := ReverseListNode.Reverse(left.Next)
	for rightHalf != nil {
		if rightHalf.Val != head.Val {
			return false
		}
		rightHalf = rightHalf.Next
		head = head.Next
	}
	return true
}
