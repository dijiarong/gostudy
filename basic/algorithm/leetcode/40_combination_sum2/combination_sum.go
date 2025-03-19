package combination_sum

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
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
			// 剪枝逻辑，值相同的树枝，只遍历第一条
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			track = append(track, candidates[i])
			backtrack(i+1, target-candidates[i])
			track = track[:len(track)-1]
		}
	}
	backtrack(0, target)
	return res
}
