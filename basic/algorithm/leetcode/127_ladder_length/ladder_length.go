package ladder_length

// bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := map[string]bool{}
	for i := range wordList {
		dict[wordList[i]] = true
	}
	if !dict[endWord] {
		return 0
	}
	q := []string{beginWord}
	visited := map[string]bool{
		beginWord: true,
	}
	res := 1
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			for _, w := range getNexts(cur, dict) {
				if w == endWord {
					return res + 1
				}
				if !visited[w] {
					q = append(q, w)
					visited[w] = true
				}
			}
		}
		res++
	}
	return 0
}

func getNexts(word string, dict map[string]bool) []string {
	res := []string{}
	tmp := []byte(word)
	for i := 0; i < len(tmp); i++ {
		old := tmp[i]
		for j := byte('a'); j <= 'z'; j++ {
			if j == old {
				continue
			}
			tmp[i] = j
			curStr := string(tmp)
			if dict[curStr] {
				res = append(res, curStr)
			}
		}
		tmp[i] = old
	}
	return res
}
