package shortestpathinbinarymatrix

// 1091. 二进制矩阵中的最短路径
// https://leetcode.cn/problems/shortest-path-in-binary-matrix/

// 宽度优先遍历
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	dir := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {1, -1}, {1, 1}, {-1, -1}, {-1, 1}}
	// 每一步可能出现的节点入库
	q := [][]int{{0, 0}}
	ans := 1
	used := map[[2]int]bool{}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, i := range tmp {
			x, y := i[0], i[1]
			if x == n-1 && y == n-1 {
				return ans
			}
			used[[2]int{x, y}] = true
			for _, d := range dir {
				dx, dy := x+d[0], y+d[1]
				if dx >= 0 && dx < n && dy >= 0 && dy < n && grid[dx][dy] == 0 && !used[[2]int{dx, dy}] {
					q = append(q, []int{dx, dy})
					used[[2]int{dx, dy}] = true
				}
			}
		}
		ans += 1
	}
	return -1
}
