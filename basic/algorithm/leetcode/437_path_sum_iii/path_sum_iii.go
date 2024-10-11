package path_sum_iii

// url:https://leetcode.cn/problems/path-sum-iii/?envType=study-plan-v2&envId=leetcode-75
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	cnt := map[int]int{
		0: 1,
	}
	var digui func(root *TreeNode, targetSum int)
	digui = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		targetSum += root.Val
		ans += cnt[targetSum]
		cnt[targetSum]++
		digui(root.Left, targetSum)
		digui(root.Right, targetSum)
		cnt[targetSum]--
	}
	return
}
