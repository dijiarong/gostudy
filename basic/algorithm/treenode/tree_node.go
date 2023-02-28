package treenode

import (
	"math"
)

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

func isSymmetric(root *TreeNode) bool {
	return true
}

func (TreeNode) isValidBST1(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}
