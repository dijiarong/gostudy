package max_sub_array

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := range nums {
		dp[i+1] = max(nums[i], dp[i]+nums[i])
	}
	res := 0
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
