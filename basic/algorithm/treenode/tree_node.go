package treenode

import (
	"math"
	"myself/mymath"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历
func (t TreeNode) FirstRange(root *TreeNode) {
	println(root.Val)
	if root.Left != nil {
		t.FirstRange(root.Left)
	}
	if root.Right != nil {
		t.FirstRange(root.Right)
	}
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

func (TreeNode) isAVLTree(root *TreeNode) bool {
	res, _ := helper1(root)
	return res
}

func helper1(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lRes, lh := helper1(root.Left)
	rRes, rh := helper1(root.Right)
	return lRes && rRes && ((lh-rh) >= -1 && (lh-rh) <= 1), mymath.Math.GetMaxInt(lh, rh) + 1
}
