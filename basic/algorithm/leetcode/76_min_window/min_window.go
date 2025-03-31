package min_window

func minWindow(s string, t string) string {
	target, window := map[byte]int{}, map[byte]int{}
	for i := range t {
		target[t[i]]++
	}
	invalid := 0
	l, r := 0, 0
	res := s + " "
	for r < len(s) {
		c := s[r]
		r++
		if tmp, ok := target[c]; ok {
			window[c]++
			if window[c] == tmp {
				invalid++
			}
		}
		// 缩小窗口
		for invalid == len(target) {
			if r-l <= len(res) {
				res = s[l:r]
			}
			c = s[l]
			l++
			if tmp, ok := target[c]; ok {
				if window[c] == tmp {
					invalid--
				}
				window[c]--
			}
		}
	}
	if res == s+" " {
		return ""
	}
	return res
}
