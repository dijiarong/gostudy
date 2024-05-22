package dailytemperatures

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/daily-temperature

// dailyTemperatures 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，
// 返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。
func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Print(dailyTemperatures(temperatures))
	fmt.Print("\n")
}
