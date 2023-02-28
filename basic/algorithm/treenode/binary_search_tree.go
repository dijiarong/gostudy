package treenode

import "math"

// 二叉搜索树：对于任意一非空节点，其左子树上所有的节点的值都小于当前节点， 右子树的值都大于当前节点

// 力扣：
func IsBinarySearchTree(root *TreeNode) bool {
	return IsBinarySearchTreeReq(root, math.MinInt, math.MaxInt)
}

// 方法1
func IsBinarySearchTreeReq(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return IsBinarySearchTreeReq(root.Left, min, root.Val) && IsBinarySearchTreeReq(root.Right, root.Val, max)
}
