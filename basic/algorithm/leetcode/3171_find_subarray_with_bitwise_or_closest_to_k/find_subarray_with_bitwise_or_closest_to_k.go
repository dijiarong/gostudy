package find_subarray_with_bitwise_or_closest_to_k

import "math"

func minimumDifference(nums []int, k int) int {
	res := math.MaxInt
	for i, x := range nums {
		res = min(x-k, res)
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			res = min(nums[j]-k, res)
			if res == 0 {
				return 0
			}
		}
	}
	return res
}

// 获取最小值
func min(a, b int) int {
	if a < 0 {
		a = -a
	}
	if a < b {
		return a
	}
	return b
}
