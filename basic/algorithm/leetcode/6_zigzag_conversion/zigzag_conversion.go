package zigzag_conversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	levelStr := make([]string, numRows)
	isTurn := -1
	curIndex := 0
	for _, cur := range s {
		levelStr[curIndex] += string(cur)
		if curIndex == 0 || curIndex == numRows-1 {
			isTurn = -isTurn
		}
		curIndex += isTurn
	}
	return strings.Join(levelStr, "")
}
