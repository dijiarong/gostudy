package permute_unique

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	track := []int{}
	used := make([]bool, n)
	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	var backtrack func()
	backtrack = func() {
		if len(track) == n {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtrack()
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack()
	return res
}
