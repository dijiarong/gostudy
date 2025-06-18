package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}
	tmp := []int{}
	for x != 0 {
		tmp = append(tmp, x%10)
		x /= 10
	}
	for i := 0; i < len(tmp)/2; i++ {
		if tmp[i] != tmp[len(tmp)-1-i] {
			return false
		}
	}
	return true
}
