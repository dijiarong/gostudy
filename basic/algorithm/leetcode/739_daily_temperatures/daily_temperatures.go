package dailytemperatures

// https://leetcode.cn/problems/daily-temperature
// 使用单调栈
// dailyTemperatures 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，
// 返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。
func dailyTemperatures(temperatures []int) []int {
	// 定义一个单调栈
	stack := make([]int, len(temperatures))
	// 定义栈顶指针
	top := -1
	tmp := make([]int, len(temperatures))
	for i := range temperatures {
		// 不断判断当前温度是否大于栈顶元素，大于时，栈顶元素出栈，并记录下当前元素与栈顶元素之间的距离
		for top != -1 && temperatures[i] > temperatures[stack[top]] {
			tmp[stack[top]] = i - stack[top]
			top--
		}
		top++
		stack[top] = i
	}
	return tmp
}
