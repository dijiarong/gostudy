package postordertraversal

// 145.二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	var process func(node *TreeNode)
	process = func(node *TreeNode) {
		if node == nil {
			return
		}
		process(node.Left)
		process(node.Right)
		res = append(res, node.Val)
	}
	process(root)
	return res
}

// 迭代
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{root}
	used := map[*TreeNode]bool{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if used[node] {
			res = append(res, node.Val)
		} else {
			stack = append(stack, node)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			used[node] = true
		}
	}
	return res
}
