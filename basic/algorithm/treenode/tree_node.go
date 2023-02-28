package treenode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历
func (t TreeNode) Range(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val) // 先序遍历
	t.Range(root.Left)
	println(root.Val) // 中序遍历
	t.Range(root.Right)
	println(root.Val) // 后序遍历
}
