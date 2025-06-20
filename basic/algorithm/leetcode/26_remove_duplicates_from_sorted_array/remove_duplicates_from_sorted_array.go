package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	// 双指针
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}
