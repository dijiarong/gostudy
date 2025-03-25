package maximum_or

func maximumOr(nums []int, k int) int {
	n := len(nums)
	orSum := 0
	for _, v := range nums {
		orSum |= v
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = orSum
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = max(nums[i-1]*2|dp[i-1][j-1], dp[i][j-1])
		}
	}
	return dp[n][k]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
