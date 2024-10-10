package sub_array_sum_equals_k

func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	for i := range nums {
		sum[i+1] = sum[i] + nums[i]
	}
	ans := 0
	countMap := make(map[int]int, len(nums))
	for _, v := range sum {
		ans += countMap[v-k]
		countMap[v]++
	}
	return ans
}
