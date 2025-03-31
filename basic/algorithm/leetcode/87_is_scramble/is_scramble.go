package is_scramble

func isScramble(s1 string, s2 string) bool {
	// dp
	m, n := len(s1), len(s2)
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	// dp[i][j][l]: 代表从 s1从i 开始长度l 和s2从j开始长度l是否为扰动字符串如果 1是 0否 -1未进行查看
	var dfs func(s1, s2 string, i1, i2, l int) bool
	dfs = func(s1, s2 string, i1, i2, l int) bool {
		// 判断是否相同
		if s1[i1:i1+l] == s2[i2:i2+l] {
			dp[i1][i2][l] = 1
			return true
		}
		// 判断两边对应的字符数量是否匹配
		if !isValid(s1, s2, i1, i2, l) {
			dp[i1][i2][l] = 0
			return false
		}
		if dp[i1][i2][l] != -1 {
			return dp[i1][i2][l] == 1
		}
		for i := 1; i < l; i++ {
			// 左右子串不进行交换
			if dfs(s1, s2, i1, i2, i) && dfs(s1, s2, i1+i, i2+i, l-i) {
				dp[i1][i2][l] = 1
				return true
			}
			// 交换
			if dfs(s1, s2, i1, i2+l-i, i) && dfs(s1, s2, i1+i, i2, l-i) {
				dp[i1][i2][l] = 1
				return true
			}
		}
		dp[i1][i2][l] = 0
		return false
	}
	return dfs(s1, s2, 0, 0, n)
}

func isValid(s1, s2 string, i1, i2, l int) bool {
	tmp := map[byte]int{}
	for i := i1; i < l; i++ {
		tmp[s1[i]]++
	}
	for i := i2; i < l; i++ {
		tmp[s2[i]]--
		if tmp[s2[i]] < 0 {
			return false
		}
		if tmp[s2[i]] == 0 {
			delete(tmp, s2[i])
		}
	}
	return len(tmp) == 0
}
