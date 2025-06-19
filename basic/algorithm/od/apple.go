package od

// 1. A和B两个人分到的苹果重量进行二进制不进位加法结果要相等.
// 2. 在满足1的情况下,B分到苹果重量正常进行十进制加法最大的一种情况
func maxWeightForB(weight []int) int {
	// 统计数据
	n := len(weight)
	if n == 0 {
		return -1
	}
	// 统计totalXor
	totalXor := 0
	totalWeight := 0
	for _, w := range weight {
		totalXor ^= w
		totalWeight += w
	}
	// 需要异或结果相同的两组，如果整体xor不等于0则证明不可平分
	if totalXor != 0 {
		return -1
	}
	dp := map[int]int{
		0: 0,
	}
	for _, w := range weight {
		tmpDP := map[int]int{}
		// 不选择当前苹果
		for xor, sumWeight := range dp {
			tmpDP[xor] = max(sumWeight, tmpDP[xor])
		}
		// 选择当前苹果
		for xor, sumWeight := range dp {
			tmpXor := xor ^ w
			tmpDP[tmpXor] = max(tmpDP[tmpXor], sumWeight+w)
		}
		dp = tmpDP
	}
	res := 0
	for _, w := range dp {
		if w > 0 && w < totalWeight {
			tmp := totalWeight - w
			res = max(res, tmp)
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
