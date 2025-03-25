package max_profit

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i, v := range prices {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -v
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+v)
		dp[i][1] = max(dp[i-1][1], -v)
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
