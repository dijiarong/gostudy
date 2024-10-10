package find_the_number_of_good_pairs_i

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	var ans int64 = 0
	tmp := map[int]int{}
	for _, v := range nums1 {
		if v%k != 0 {
			continue
		}
		v /= k
		// 记录v/k的因子
		for i := 1; i*i <= v; i++ {
			if v%i == 0 {
				tmp[i]++
				if i*i != v {
					tmp[v/i]++
				}
			}
		}
	}
	for _, v := range nums2 {
		ans += int64(tmp[v])
	}
	return ans
}
