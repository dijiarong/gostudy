package divide_array_into_arrays_with_max_difference

import (
	"slices"
)

func divideArray(nums []int, k int) [][]int {
	n := len(nums) / 3
	res := make([][]int, 0, n)
	slices.Sort(nums)
	for i := 0; i < n; i++ {
		if nums[i*3+2]-nums[i*3] > k {
			return nil
		}
		res = append(res, nums[i*3:i*3+3])
	}
	return res
}
