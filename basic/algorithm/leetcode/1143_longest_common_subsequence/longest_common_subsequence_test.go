package longest_common_subsequence

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
}
