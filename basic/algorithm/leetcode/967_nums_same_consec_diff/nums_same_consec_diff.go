package nums_same_consec_diff

// 回溯 然后剪枝
func numsSameConsecDiff(n, k int) []int {
	res := []int{}
	track := 0
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == n {
			res = append(res, track)
			return
		}
		for i := 0; i <= 9; i++ {
			if start == 0 && i == 0 {
				continue
			}
			if start > 0 && abs(track%10-i) != k {
				continue
			}
			track = track*10 + i
			backtrack(start + 1)
			track = track / 10
		}
	}
	backtrack(0)
	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
