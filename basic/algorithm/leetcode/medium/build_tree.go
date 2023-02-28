package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 105：从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) {
		return nil
	}
	return buildTreeReq(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTreeReq(
	preorder []int, inorder []int,
	l1, r1, l2, r2 int,
) *TreeNode {
	if l1 > r1 {
		return nil
	}
	node := &TreeNode{
		Val:   preorder[l1],
		Left:  nil,
		Right: nil,
	}
	if l1 == r1 {
		return node
	}
	// 找到头节点
	find := l2
	for preorder[l1] != inorder[find] {
		find++
	}
	node.Left = buildTreeReq(preorder, inorder, l1+1, l1+find-l2, l2, find-1)
	node.Right = buildTreeReq(preorder, inorder, l1+find-l2+1, r1, find+1, r2)
	return node
}

func RemoveError() {
	buildTree(nil, nil)
}
