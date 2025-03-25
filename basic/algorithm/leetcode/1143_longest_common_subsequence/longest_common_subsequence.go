package longest_common_subsequence

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dpTable := make([][]int, l1+1)
	for i := range dpTable {
		dpTable[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dpTable[i][j] = dpTable[i-1][j-1] + 1
			} else {
				dpTable[i][j] = max(dpTable[i-1][j], dpTable[i][j-1], dpTable[i-1][j-1])
			}
		}
	}
	return dpTable[l1][l2]
}

func max(args ...int) int {
	res := math.MinInt
	for _, v := range args {
		if res < v {
			res = v
		}
	}
	return res
}
