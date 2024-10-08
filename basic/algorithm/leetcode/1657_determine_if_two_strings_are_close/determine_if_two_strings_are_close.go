package determine_if_two_strings_are_close

import (
	"reflect"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	bytes1 := make([]int, 26)
	for _, v := range word1 {
		bytes1[v-'a']++
	}
	bytes2 := make([]int, 26)
	for _, v := range word2 {
		bytes2[v-'a']++
	}
	for i := 0; i < 26; i++ {
		if bytes1[i] == 0 && bytes2[i] != 0 || bytes1[i] != 0 && bytes2[i] == 0 {
			return false
		}
	}
	sort.Slice(bytes1, func(i, j int) bool {
		return bytes1[i] >= bytes1[j]
	})
	sort.Slice(bytes2, func(i, j int) bool {
		return bytes2[i] >= bytes2[j]
	})
	return reflect.DeepEqual(bytes1, bytes2)
}
