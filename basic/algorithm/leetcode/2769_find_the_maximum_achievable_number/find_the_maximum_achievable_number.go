package findthemaximumachievablenumber

// 地址：https://leetcode.cn/problems/find-the-maximum-achievable-number
// easy
func theMaximumAchievableX(num int, t int) int {
	return num + t<<1
}
