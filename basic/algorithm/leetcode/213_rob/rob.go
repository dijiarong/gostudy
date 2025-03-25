package robs

func rob(nums []int) int {
	return max(robWithStartEnd(nums, 0, len(nums)-2), robWithStartEnd(nums, 1, len(nums)-1))
}

func robWithStartEnd(nums []int, start, end int) int {
	i, j := 0, nums[start]
	for start := start + 1; start <= end; start++ {
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
