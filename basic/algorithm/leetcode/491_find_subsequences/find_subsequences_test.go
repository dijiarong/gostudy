package find_subsequences

import (
	"fmt"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
}
