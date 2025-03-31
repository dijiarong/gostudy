package maximal_rectangle

func maximalRectangle(matrix [][]byte) int {
	res := 0
	row, col := len(matrix), len(matrix[0])
	height := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		res = max(res, maxA(height))
	}
	return res
}

func maxA(heights []int) int {
	res := 0
	stack := make([]int, len(heights))
	index := -1
	for i := 0; i < len(heights); i++ {
		for index != -1 && heights[i] <= heights[stack[index]] {
			cur := stack[index]
			index--
			l := i
			if index != -1 {
				l = l - stack[index] - 1
			}
			res = max(res, l*heights[cur])
		}
		index++
		stack[index] = i
	}
	for index != -1 {
		cur := stack[index]
		index--
		l := len(heights)
		if index != -1 {
			l = l - stack[index] - 1
		}
		res = max(res, l*heights[cur])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
