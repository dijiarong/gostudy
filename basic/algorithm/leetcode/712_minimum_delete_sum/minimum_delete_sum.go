package minimum_delete_sum

func minimumDeleteSum(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	memo := make([][]int, l1)
	for i := range memo {
		memo[i] = make([]int, l2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dp func(l, r int) int
	dp = func(l, r int) int {
		res := 0
		if l == l1 {
			for i := r; i < l2; i++ {
				res += int(s2[i])
			}
			return res
		}
		if r == l2 {
			for i := l; i < l1; i++ {
				res += int(s1[i])
			}
			return res
		}
		if memo[l][r] != -1 {
			return memo[l][r]
		}
		if s1[l] == s2[r] {
			memo[l][r] = dp(l+1, r+1)
		} else {
			memo[l][r] = min(
				int(s1[l])+dp(l+1, r),
				int(s2[r])+dp(l, r+1),
			)
		}
		return memo[l][r]
	}
	return dp(0, 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
