package find_substring

// 滑动窗口
func findSubstring(s string, words []string) (ans []int) {
	ls, wc, lw := len(s), len(words), len(words[0])
	for i := 0; i < lw && i+lw*wc <= ls; i++ {
		windows := map[string]int{}
		for j := 0; j < wc; j++ {
			windows[s[j*lw+i:(j+1)*lw+i]]++
		}
		for _, v := range words {
			windows[v]--
			if windows[v] == 0 {
				delete(windows, v)
			}
		}
		if len(windows) == 0 {
			ans = append(ans, i)
		}
		// 需要继续移动窗口直到所有的word都进行查询
		for j := i + lw; j+lw*wc <= ls; j += lw {
			// 获取进出窗口的
			sub := s[j+lw*(wc-1) : j+lw*wc]
			windows[sub]++
			if windows[sub] == 0 {
				delete(windows, sub)
			}
			// 进入一个就需要弹出一个
			sub = s[j-lw : j]
			windows[sub]--
			if windows[sub] == 0 {
				delete(windows, sub)
			}
			if len(windows) == 0 {
				ans = append(ans, j)
			}
		}
	}
	return
}
