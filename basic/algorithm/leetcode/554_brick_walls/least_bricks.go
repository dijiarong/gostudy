package brickwalls

// 554.砖墙
// https://leetcode.cn/problems/brick-wall/description/
func leastBricks(wall [][]int) int {
	b := map[int]int{}
	for _, w := range wall {
		sum := 0
		for i := 0; i < len(w)-1; i++ {
			sum += w[i]
			b[sum]++
		}
	}
	max := 0
	for _, v := range b {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}
