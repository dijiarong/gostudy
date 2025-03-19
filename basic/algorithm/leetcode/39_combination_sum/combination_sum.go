package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	track := []int{}
	var backtrack func(start, target int)
	backtrack = func(start, target int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			track = append(track, candidates[i])
			backtrack(i, target-candidates[i])
			track = track[:len(track)-1]
		}
	}
	backtrack(0, target)
	return res
}
