package find_the_xor_of_numbers_which_appear_twice

func duplicateNumbersXOR(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	ans := 0
	tmp := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := tmp[v]; ok {
			ans = ans ^ v
			continue
		}
		tmp[v] = struct{}{}
	}
	return ans ^ 0
}
