package get_permutation

import "fmt"

func getPermutation(n int, k int) string {
	var backtrack func()
	used := make([]bool, n+1)
	path := []int{}
	backtrack = func() {
		if len(path) == n {
			k--
			return
		}
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			path = append(path, i)
			used[i] = true
			backtrack()
			if k == 0 {
				return
			}
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	res := ""
	for _, v := range path {
		res += fmt.Sprintf("%d", v)
	}
	return res
}
