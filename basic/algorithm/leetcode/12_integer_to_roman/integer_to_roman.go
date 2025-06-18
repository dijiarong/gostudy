package integer_to_roman

import "strings"

var bitMap = map[int]map[int]string{
	0: map[int]string{
		1: "I",
		4: "IV",
		9: "IX",
		5: "V",
	},
	1: map[int]string{
		1: "X",
		4: "XL",
		9: "XC",
		5: "L",
	},
	2: map[int]string{
		1: "C",
		4: "CD",
		9: "CM",
		5: "D",
	},
	3: map[int]string{
		1: "M",
	},
}

func intToRoman(num int) string {
	bit := 0
	res := ""
	for num != 0 {
		cur := num % 10
		if t, ok := bitMap[bit][cur]; ok {
			res = t + res
		} else {
			c := ""
			if cur > 5 {
				c += bitMap[bit][5]
				cur -= 5
			}
			if cur != 0 {
				c += strings.Repeat(bitMap[bit][1], cur)
			}
			res = c + res
		}
		bit++
		num /= 10
	}
	return res
}
