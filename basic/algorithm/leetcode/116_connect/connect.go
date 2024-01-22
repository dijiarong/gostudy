package connect

// 116.填充每个节点的下一个右侧节点指针
// 题：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	// 层序遍历
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i := 0; i < len(tmp); i++ {
			if i != 0 {
				tmp[i-1].Next = tmp[i]
			}
			if tmp[i].Left != nil {
				q = append(q, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				q = append(q, tmp[i].Right)
			}
		}
	}
	return root
}
