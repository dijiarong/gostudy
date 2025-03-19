package word_break

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
