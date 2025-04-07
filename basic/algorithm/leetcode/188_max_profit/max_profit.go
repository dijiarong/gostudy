package max_profit

import "math"

func maxProfit(k int, prices []int) int {
	// 定义dp[day][k][2]
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
		dp[i][0][1] = math.MinInt
	}
	// 初始化边界
	dp[0][k][1] = -prices[0]
	for i := 0; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[n][k][0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
