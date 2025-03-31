package largestrectangleinhistogram

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	// heights := []int{2, 1, 5, 6, 2, 3}
	heights := []int{3, 6, 5, 7, 4, 8, 1, 0}
	fmt.Println(LargestRectangleArea(heights))
}
