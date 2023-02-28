package listnode

type cycleListNode struct{}

var CycleListNode cycleListNode

// 笨办法，循环超过链表最大长度即为true，中途出现了nil则证明无问题
// 链表长度较短的时候，会浪费性能
func (cycleListNode) HasCycle1(head *ListNode) bool {
	i := 0
	for i <= 10000 {
		if head == nil {
			return false
		}
		head = head.Next
		i++
	}
	return true
}

// 记录当前的值，循环后面的是否相等
func (cycleListNode) HasCycle2(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
