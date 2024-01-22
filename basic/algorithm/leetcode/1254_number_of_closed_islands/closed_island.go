package numberofclosedislands

// 1254. 统计封闭岛屿的数目
// https://leetcode.cn/problems/number-of-closed-islands/

// 宽度优先遍历，遍历到的每个0都需要把自己的上下左右是0的加入队列
func closedIsland(grid [][]int) int {
	dir := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				q := [][]int{{i, j}}
				closed := true
				for len(q) > 0 {
					node := q[0]
					q = q[1:]
					x, y := node[0], node[1]
					if x == 0 || x == m-1 || y == 0 || y == n-1 {
						closed = false
					}
					for _, d := range dir {
						dx, dy := x+d[0], y+d[1]
						if dx >= 0 && dx <= m-1 && dy >= 0 && dy <= n-1 && grid[dx][dy] == 0 {
							grid[dx][dy] = 1
							q = append(q, []int{dx, dy})
						}
					}
				}
				if closed {
					ans += 1
				}
			}
		}
	}
	return ans
}

// 深度优先遍历,遍历到的每个为0的位置时，需要获取其上下左右的信息
func closedIsland1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		if grid[x][y] != 0 {
			return true
		}
		grid[x][y] = -1
		ret1, ret2, ret3, ret4 := dfs(x+1, y), dfs(x-1, y), dfs(x, y+1), dfs(x, y-1)
		return ret1 && ret2 && ret3 && ret4
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && dfs(i, j) {
				ans += 1
			}
		}
	}
	return ans
}
