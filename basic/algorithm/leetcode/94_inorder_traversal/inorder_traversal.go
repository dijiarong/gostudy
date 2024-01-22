package inordertraversal

// 94.二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方法
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	var process func(node *TreeNode)
	process = func(node *TreeNode) {
		if root == nil {
			return
		}
		process(root.Left)
		res = append(res, root.Val)
		process(root.Right)
	}
	process(root)
	return res
}

// 迭代方法,把所有的左子树入栈，当出栈的时候加入res
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
