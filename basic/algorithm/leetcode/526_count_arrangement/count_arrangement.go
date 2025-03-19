package count_arrangement

func countArrangement(n int) int {
	res := 0
	used := make([]bool, n+1)
	var backtrack func(i int)
	backtrack = func(start int) {
		if start > n {
			res++
			return
		}

		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			if i%start != 0 && start%i != 0 {
				continue
			}
			used[i] = true
			backtrack(start + 1)
			used[i] = false
		}
	}
	backtrack(1)
	return res
}
