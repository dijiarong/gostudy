package can_partition_k_subsets

func canPartitionKSubsets(nums []int, k int) bool {
	target := 0
	for _, v := range nums {
		target += v
	}
	if target%k != 0 {
		return false
	}
	target /= k
	track := make([]int, k)
	// 一个一个子集进行处理
	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		// 当查到
		if index == len(nums) {
			for _, v := range track {
				if v != target {
					return false
				}
			}
			return true
		}
		for i := 0; i < len(track); i++ {
			if track[i]+nums[index] > target {
				continue
			}
			track[i] += nums[index]
			if backtrack(index + 1) {
				return true
			}
			track[i] -= nums[index]
		}
		return false
	}
	return backtrack(0)
}
