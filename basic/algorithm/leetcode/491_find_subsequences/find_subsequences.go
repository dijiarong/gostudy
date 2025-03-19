package find_subsequences

func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(track) > 1 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
		}
		used := map[int]bool{}
		for i := index; i < len(nums); i++ {
			if len(track) > 0 && nums[i] < track[len(track)-1] {
				continue
			}
			// 本层不能选择已选择的元素
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
