package full_justify

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	curWords := []string{}
	curWordLen := 0
	for _, v := range words {
		if curWordLen+len(curWords)+len(v) <= maxWidth {
			curWords = append(curWords, v)
			curWordLen += len(v)
		} else {
			res = append(res, deal(curWords, curWordLen, maxWidth))
			curWords = []string{v}
			curWordLen = len(v)
		}
	}
	// 处理最后一行
	last := strings.Join(curWords, " ")
	if len(last) < maxWidth {
		last += strings.Repeat(" ", maxWidth-len(last))
	}
	return append(res, last)
}

func deal(words []string, curWordLen, maxWidth int) string {
	if len(words) == 0 {
		return ""
	}
	spaceCount := maxWidth - curWordLen
	spaceNum := []int{spaceCount}
	if len(words) > 1 {
		spaceNum = make([]int, len(words)-1)
		// 循环给所有的组增加
		for i := 0; i < spaceCount; i++ {
			spaceNum[i%len(spaceNum)]++
		}
	}
	spaceNum = append(spaceNum, 0)
	sbuilder := strings.Builder{}
	for i := range words {
		sbuilder.WriteString(words[i])
		sbuilder.WriteString(strings.Repeat(" ", spaceNum[i]))
	}
	return sbuilder.String()
}
