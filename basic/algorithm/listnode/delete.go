package listnode

type deleteListNode struct{}

var DeleteListNode deleteListNode

// 删除某个节点
func (deleteListNode) DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
