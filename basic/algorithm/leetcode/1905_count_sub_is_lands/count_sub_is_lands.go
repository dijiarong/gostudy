package count_sub_is_lands

// 解题思路
// 当grid2中为陆地而grid1中为海洋时，需要将岛屿淹没
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i, v := range grid2 {
		for j, val := range v {
			if val == 1 && grid1[i][j] == 0 {
				dfs(grid2, i, j)
			}
		}
	}
	for i, v := range grid2 {
		for j, val := range v {
			if val == 1 {
				res++
				dfs(grid2, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
}
