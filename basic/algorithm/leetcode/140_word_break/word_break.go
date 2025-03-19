package word_break

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	memo := make([][]string, len(s))
	var dp func(i int) []string
	dp = func(i int) []string {
		res := []string{}
		if i == len(s) {
			res = append(res, "")
			return res
		}
		if len(memo[i]) > 0 {
			return memo[i]
		}
		for j := 1; i+j <= len(s); j++ {
			tmp := s[i : i+j]
			if wordMap[tmp] {
				sub := dp(i + j)
				for _, v := range sub {
					if v == "" {
						res = append(res, tmp)
					} else {
						res = append(res, tmp+" "+v)
					}
				}
			}
		}
		memo[i] = res
		return res
	}
	dp(0)
	return memo[0]
}
