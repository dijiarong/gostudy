package treenode

type maxDepthTreeNode struct{}

var MaxDepth maxDepthTreeNode

// 递归法
func (m maxDepthTreeNode) MaxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMax(m.MaxDepth1(root.Left), m.MaxDepth1(root.Right)) + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
