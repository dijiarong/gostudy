package robs

func rob(nums []int) int {
	n := len(nums)
	i, j := 0, nums[0]
	for start := 1; start < n; start++ {
		tmp := max(i+nums[start], j)
		i = j
		j = tmp
	}
	return j
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
