package sliding_puzzle

import (
	"fmt"
	"testing"
)

func TestSlidingPuzzle(t *testing.T) {
	fmt.Println(slidingPuzzle([][]int{
		{1, 2, 3},
		{4, 0, 5},
	}))
	fmt.Println(slidingPuzzle([][]int{
		{4, 1, 2},
		{5, 0, 3},
	}))
}
