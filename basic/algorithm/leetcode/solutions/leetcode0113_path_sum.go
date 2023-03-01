package solutions

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	pathSumReq(root, targetSum, []int{}, 0, &res)
	return res
}

func pathSumReq(root *TreeNode, targetSum int,
	preList []int, preSum int,
	res *[][]int) {
	if root.Left == nil && root.Right == nil {
		if root.Val+preSum == targetSum {
			cp := make([]int, len(preList)+1)
			copy(cp, append(preList, root.Val))
			*res = append(*res, cp)
		}
		return
	}
	preSum += root.Val
	if root.Left != nil {
		pathSumReq(root.Left, targetSum, append(preList, root.Val), preSum, res)
	}
	if root.Right != nil {
		pathSumReq(root.Right, targetSum, append(preList, root.Val), preSum, res)
	}
}
