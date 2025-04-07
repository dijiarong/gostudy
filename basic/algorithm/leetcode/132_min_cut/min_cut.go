package min_cut

import "math"

func minCut(s string) int {
	n := len(s)
	dp := make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var isValid func(i, j int) bool
	isValid = func(i, j int) bool {
		if i >= j {
			return true
		}
		if dp[i][j] != -1 {
			return dp[i][j] == 1
		}
		res := s[i] == s[j] && isValid(i+1, j-1)
		if res {
			dp[i][j] = 1
		} else {
			dp[i][j] = 0
		}
		return res
	}
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	// 递归
	var dfs func(i int) int
	dfs = func(i int) int {
		if isValid(0, i) {
			return 0
		}
		// 记忆
		if memo[i] != -1 {
			return memo[i]
		}
		res := math.MaxInt
		for l := 1; l <= i; l++ {
			if isValid(l, i) {
				res = min(res, dfs(l-1)+1)
			}
		}
		memo[i] = res
		return res
	}
	return dfs(n - 1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
