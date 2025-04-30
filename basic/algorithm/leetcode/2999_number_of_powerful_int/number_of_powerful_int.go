package number_of_powerful_int

import (
	"fmt"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	var count int64
	// 遍历区间 [start, finish] 内的所有整数
	for num := start; num <= finish; num++ {
		numStr := fmt.Sprintf("%d", num)
		// 检查 num 的末尾部分是否是 s
		if len(numStr) >= len(s) && numStr[len(numStr)-len(s):] == s {
			isPowerful := true
			// 检查 num 中的每个数位是否至多是 limit
			for _, digit := range numStr {
				digitValue := int(digit - '0')
				if digitValue > limit {
					isPowerful = false
					break
				}
			}
			if isPowerful {
				count++
			}
		}
	}
	return count
}

func getBitNun(num int64) int {
	res := 0
	for num != 0 {
		num >>= 1
		res++
	}
	return res
}
