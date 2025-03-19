package word_break

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	memo := make([]int, len(s))
	var dp func(i int) bool
	dp = func(i int) bool {
		if i == len(s) {
			return true
		}
		if memo[i] != 0 {
			return memo[i] == 1
		}
		for length := 0; i+length <= len(s); length++ {
			tmp := s[i : i+length]
			if wordMap[tmp] {
				if dp(i + length) {
					memo[i] = 1
					return true
				}
			}
		}
		memo[i] = -1
		return false
	}
	return dp(0)
}
