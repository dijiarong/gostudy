package num_distinct

func numDistinct(s string, t string) int {
	sl, tl := len(s), len(t)
	memo := make([][]int, sl)
	for i := range memo {
		memo[i] = make([]int, tl)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dp(s, t, 0, 0, memo)
}

func dp(s, t string, sIndex, tIndex int, memo [][]int) int {
	// base case 1
	if tIndex == len(t) {
		return 1
	}
	// base case 2
	if len(s)-sIndex < len(t)-tIndex {
		return 0
	}
	if memo[sIndex][tIndex] != -1 {
		return memo[sIndex][tIndex]
	}
	res := 0
	if s[sIndex] == t[tIndex] {
		res = dp(s, t, sIndex+1, tIndex+1, memo) + dp(s, t, sIndex+1, tIndex, memo)
	} else {
		res = dp(s, t, sIndex+1, tIndex, memo)
	}
	memo[sIndex][tIndex] = res
	return res
}
