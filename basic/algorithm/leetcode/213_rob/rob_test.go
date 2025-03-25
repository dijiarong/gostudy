package robs

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
