package longest_palindromic_substring

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		a := isPalindrome(s, i, i)
		b := isPalindrome(s, i, i+1)
		if len(a) > len(res) {
			res = a
		}
		if len(b) > len(res) {
			res = b
		}
	}
	return res
}

// 以某某某为中点向外扩散最长回文串是多少
func isPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		r++
		l--
	}
	return s[l+1 : r]
}
