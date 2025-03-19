package permute

func permute(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	track := []int{}
	used := map[int]bool{}
	var backtrack func()
	backtrack = func() {
		if len(track) == n {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[nums[i]] {
				continue
			}
			track = append(track, nums[i])
			used[nums[i]] = true
			backtrack()
			track = track[:len(track)-1]
			used[nums[i]] = false
		}
	}
	backtrack()
	return res
}
