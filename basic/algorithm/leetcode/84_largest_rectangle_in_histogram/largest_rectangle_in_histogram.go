package largestrectangleinhistogram

import "math"

func LargestRectangleArea(heights []int) int {
	// 单调栈,栈内元素是heights索引信息，以heights[i]从小到大
	stack := make([]int, len(heights))
	// 单调栈栈顶index
	index := -1
	res := math.MinInt
	// 当前入栈元素比栈顶元素小时，弹出计算
	for i := 0; i < len(heights); i++ {
		for index != -1 && heights[i] <= heights[stack[index]] {
			// 弹出当前栈顶元素
			curIndex := stack[index]
			index -= 1
			// 左边界-右边界
			l := i
			if index != -1 {
				l = l - stack[index] - 1
			}
			// 计算面积
			res = getMax(res, heights[curIndex]*l)
		}
		// 入栈
		index += 1
		stack[index] = i
	}
	// 栈内还有元素时
	for index != -1 {
		// 弹出计算
		// 弹出当前栈顶元素
		curIndex := stack[index]
		index -= 1
		l := len(heights)
		if index != -1 {
			l = l - stack[index] - 1
		}
		res = getMax(res, heights[curIndex]*l)
	}
	return res
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
