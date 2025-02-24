package minimum_average_of_smallest_and_largest_elements

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := math.MaxFloat64
	for i := 0; i < len(nums)/2; i++ {
		if res > float64(nums[i]+nums[len(nums)-i-1])/2 {
			res = float64(nums[i]+nums[len(nums)-i-1]) / 2
		}
	}
	return res
}
