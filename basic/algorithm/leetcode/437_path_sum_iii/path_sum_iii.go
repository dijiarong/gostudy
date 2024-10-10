package path_sum_iii

// url:https://leetcode.cn/problems/path-sum-iii/?envType=study-plan-v2&envId=leetcode-75
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	return digui(root, targetSum)
}

func digui(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return 0
	}
	if targetSum == 0 {
		res += 1
	}
	return res + digui(root.Left, targetSum-root.Val) + digui(root.Right, targetSum-root.Val) + digui(root.Left, targetSum) + digui(root.Right, targetSum)
}
