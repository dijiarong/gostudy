package unique_paths3

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	track := 0
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var startI, startJ int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startI, startJ = i, j
			}
			if grid[i][j] == 1 || grid[i][j] == 0 {
				track++
			}
		}
	}
	var backtrack func(i, j int)
	backtrack = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid[i][j] == 2 {
			if track == 0 {
				res++
			}
			return
		} else if grid[i][j] == -1 {
			return
		}
		if used[i][j] {
			return
		}
		used[i][j] = true
		track--
		for _, v := range dirs {
			backtrack(i+v[0], j+v[1])
		}
		used[i][j] = false
		track++
	}
	backtrack(startI, startJ)
	return res
}
