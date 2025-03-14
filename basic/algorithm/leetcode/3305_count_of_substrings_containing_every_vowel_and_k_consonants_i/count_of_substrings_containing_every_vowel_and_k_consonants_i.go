package count_of_substrings_containing_every_vowel_and_k_consonants_i

var yuanMap = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func countOfSubstrings(word string, k int) int {
	return f(word, k) - f(word, k+1)
}

func f(word string, k int) int {
	n := len(word)
	result := 0
	left, right := 0, 0
	cnt1 := map[byte]int{}
	cnt2 := 0
	for right < n {
		cur := word[right]
		right++
		if yuanMap[cur] {
			cnt1[cur]++
		} else {
			cnt2++
		}
		for len(cnt1) == 5 && cnt2 >= k {
			tmp := word[left]
			if yuanMap[tmp] {
				cnt1[tmp]--
				if cnt1[tmp] == 0 {
					delete(cnt1, tmp)
				}
			} else {
				cnt2--
			}
			left++
		}
		result += left
	}

	return result
}
