package letter_combinations_of_a_phone_number

var num2ByteMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	res := []string{}
	n := len(digits)
	var trackback func(preBytes []byte, index int)
	trackback = func(preBytes []byte, index int) {
		if index == n {
			res = append(res, string(preBytes))
			return
		}
		for _, cur := range num2ByteMap[digits[index]] {
			preBytes = append(preBytes, cur)
			trackback(preBytes, index+1)
			preBytes = preBytes[:len(preBytes)-1]
		}
	}
	trackback([]byte{}, 0)
	return res
}
