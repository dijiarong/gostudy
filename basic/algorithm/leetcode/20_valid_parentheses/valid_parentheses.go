package valid_parentheses

func isValid(s string) bool {
	leftMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, len(s))
	index := 0
	for i := 0; i < len(s); i++ {
		right, ok := leftMap[s[i]]
		if !ok {
			stack[index] = s[i]
			index++
		} else {
			if index > 0 && stack[index-1] == right {
				index--
			} else {
				return false
			}
		}
	}
	return index == 0
}
