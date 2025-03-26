package is_match

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	// dp[i][j] P中前J个字符是否可以匹配s的前i
	dp := make([][]bool, ls+1)
	for i := range dp {
		dp[i] = make([]bool, lp+1)
	}
	dp[0][0] = true
	for i := 1; i <= lp; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[ls][lp]
}
