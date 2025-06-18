package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	// 双指针记录两指针直接的字符出现次数
	window := map[byte]int{}
	left := 0
	res := 0
	for i := range s {
		cur := s[i]
		window[cur]++
		for window[cur] > 1 {
			tmp := s[left]
			left++
			window[tmp]--
		}
		if i-left+1 > res {
			res = i - left + 1
		}
	}
	return res
}
