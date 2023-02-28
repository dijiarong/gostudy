package treenode

// 题：https://leetcode.cn/problems/path-sum/
var res = false

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	res = false
	HasPathSumHelp(root, 0, targetSum)
	return res
}

func HasPathSumHelp(root *TreeNode, preSum, targetSum int) {
	if root.Left == nil && root.Right == nil {
		if preSum+root.Val == targetSum {
			res = true
		}
		return
	}
	preSum += root.Val
	if root.Left != nil {
		HasPathSumHelp(root.Left, preSum, targetSum)
	}
	if root.Right != nil {
		HasPathSumHelp(root.Right, preSum, targetSum)
	}
}
