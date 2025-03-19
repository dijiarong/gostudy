package max_area_of_is_land

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i, v := range grid {
		for j, val := range v {
			if val == 1 {
				res = max(res, dfs(grid, i, j))
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfs(grid, i, j+1) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i-1, j)
}
