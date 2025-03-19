package num_is_lands

func numIslands(grid [][]byte) int {
	res := 0
	for i, v := range grid {
		for j, val := range v {
			if val == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// 将传入坐标及相邻土地转成海洋
func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}
