package calculate_minimum_hp

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	// 表示从i, j 走到终点需要多少生命值
	// 只能走 i, j+1, i+1, j 两个，所以应该根据这个来判断
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	// 终点的下一个位置初始化为 1
	dp[m][n-1] = 1
	dp[m-1][n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			m := min(dp[i][j+1], dp[i+1][j])
			dp[i][j] = max(m-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
