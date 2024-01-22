package preordertraversal

// 144.二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	var process func(node *TreeNode)
	process = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		process(node.Left)
		process(node.Right)
	}
	process(root)
	return res
}

// 迭代
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
