package closed_is_land

func closedIsland(grid [][]int) int {
	res := 0
	// 将靠边的岛屿淹没
	m, n := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		dfs(grid, 0, i)
		dfs(grid, m-1, i)
	}
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}
	for i, v := range grid {
		for j, val := range v {
			if val == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
}
