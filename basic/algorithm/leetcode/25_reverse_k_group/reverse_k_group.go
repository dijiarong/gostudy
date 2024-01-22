package reversekgroup

import "fmt"

// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{0, head}
	pro := pre
	for head != nil {
		// 获取k个end
		end := getNode(head, k)
		if end == nil {
			return pro.Next
		}
		// 记录下一组的next
		nextHead := end.Next
		// 反转
		reverse(head, nextHead)
		// 两组之间的连接
		pre.Next = end
		head.Next = nextHead
		pre = head
		head = nextHead
	}
	p(pro)
	return pro.Next
}

func p(root *ListNode) {
	node := root
	res := []int{}
	for node.Next != nil {
		res = append(res, node.Next.Val)
	}
	fmt.Print(res)
}

func reverse(head, end *ListNode) {
	pre := &ListNode{}
	node := head
	for node != end {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
}

func getNode(node *ListNode, k int) *ListNode {
	for node != nil {
		k--
		if k == 0 {
			return node
		}
		node = node.Next
	}
	return nil
}
