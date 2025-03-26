package is_match

func isMatch(s string, p string) bool {
	// 动态规划
	// dp[i][j] 代表 s[:i] p[:j]是否能匹配
	// dp[i][j] =
	// 1.当p[j-1] = . || s[i-1] ==j[j-1]  dp[i][j] = dp[i-1][j-1]
	// 2.当p[j-1] = *
	// 2.1 匹配0次 dp[i][j] = dp[i][j-2]
	// 2.2 匹配多次 需要取决于s[i] 是否匹配与p[j-1]
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}
