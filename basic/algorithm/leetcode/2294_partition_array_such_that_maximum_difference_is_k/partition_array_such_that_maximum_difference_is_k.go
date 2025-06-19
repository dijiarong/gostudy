package partition_array_such_that_maximum_difference_is_k

import (
	"math"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	nextMin := math.MinInt / 2
	res := 0
	for _, num := range nums {
		if num-nextMin > k {
			res++
			nextMin = num
		}
	}
	return res
}
