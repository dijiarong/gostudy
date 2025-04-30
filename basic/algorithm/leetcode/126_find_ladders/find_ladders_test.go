package find_ladders

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	beginWord := "lost"
	endWord := "miss"
	wordList := []string{"most", "mist", "miss", "lost", "fist", "fish"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}
