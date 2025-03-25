package find_target_sum_ways

func findTargetSumWays(nums []int, target int) int {
	//  A - B = target
	// 2A = A+B +target
	// A= SUM + TARGET/ 2
	numSum := target
	for i := range nums {
		numSum += nums[i]
	}
	if numSum%2 != 0 {
		return 0
	}
	t := numSum / 2
	// i 个物品 j重量 的次数
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= t; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][t]
}
