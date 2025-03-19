package combine

func combine(n int, k int) [][]int {
	res := [][]int{}
	track := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(1)
	return res
}
