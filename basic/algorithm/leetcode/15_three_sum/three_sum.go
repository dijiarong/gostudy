package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	if nums[0] > 0 || nums[n-1] < 0 {
		return nil
	}
	res := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r, target := n-1, 0-nums[i]
		// 两数之和
		for j := i + 1; j < r; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < r && nums[j]+nums[r] > target {
				r--
			}
			if j >= r {
				break
			}
			if nums[j]+nums[r] == target {
				res = append(res, []int{nums[i], nums[j], nums[r]})
			}
		}
	}
	return res
}
