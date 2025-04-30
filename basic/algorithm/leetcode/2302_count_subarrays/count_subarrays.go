package count_subarrays

func countSubarrays(nums []int, k int64) (count int64) {
	// 滑动窗口
	var l, r, sum int64
	for r = 0; r < int64(len(nums)); r++ {
		sum += int64(nums[r])
		for l <= r && k <= (r-l+int64(1))*sum {
			sum -= int64(nums[l])
			l++
		}
		count += r - l + 1
	}
	return
}
