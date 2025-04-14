package find_ladders

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := map[string]bool{}
	for _, word := range wordList {
		dict[word] = true
	}
	if !dict[endWord] {
		return nil
	}
	res := [][]string{}
	track := []string{}
	minl := len(wordList)
	used := map[string]bool{
		beginWord: true,
	}
	var backtrack func(s string)
	backtrack = func(s string) {
		if len(track) > minl {
			return
		}
		if s == endWord {
			if len(track) > minl {
				return
			} else if len(track) < minl {
				minl = len(track)
				res = [][]string{}
			}
			// tmp := make([]string, len(track))
			// copy(tmp, track)
			res = append(res, append([]string{beginWord}, track...))
			return
		}
		for _, word := range wordList {
			if used[word] {
				continue
			}
			if !isValid(s, word) {
				continue
			}
			track = append(track, word)
			used[word] = true
			backtrack(word)
			used[word] = false
			track = track[:len(track)-1]
		}
	}
	backtrack(beginWord)

	return res
}

func isValid(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diffCount := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
	}
	return diffCount == 1
}
