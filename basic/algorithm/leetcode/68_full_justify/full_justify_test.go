package full_justify

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	max := 16
	res := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	for i, v := range fullJustify(words, max) {
		fmt.Println(res[i] == v)
	}
}
