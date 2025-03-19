package combination_sum3

func combinationSum3(k int, n int) [][]int {
	track := []int{}
	res := [][]int{}
	var backtrack func(start, target int)
	backtrack = func(start, target int) {
		if target < 0 {
			return
		}
		if target == 0 && len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i <= 9; i++ {
			track = append(track, i)
			backtrack(i+1, target-i)
			track = track[:len(track)-1]
		}
	}
	backtrack(1, n)
	return res
}
