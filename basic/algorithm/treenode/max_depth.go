package treenode

import "gostudy/basic/mymath"

type maxDepthTreeNode struct{}

var MaxDepth maxDepthTreeNode

// 递归法
func (m maxDepthTreeNode) MaxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return mymath.Math.GetMaxInt(m.MaxDepth1(root.Left), m.MaxDepth1(root.Right)) + 1
}
