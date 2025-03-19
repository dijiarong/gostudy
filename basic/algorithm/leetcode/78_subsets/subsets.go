package subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
