package shortest_palindrome

// 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。
// 找到并返回可以用这种方式转换的最短回文串
func shortestPalindrome(s string) string {
	// 感觉需要找到前缀中最大的回文串
	for i := len(s) - 1; i >= 0; i-- {
		if isPalindrome(s[:i]) {
			return reverse(s[i:]) + s
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func reverse(s string) string {
	tmp := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}
