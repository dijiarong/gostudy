package unique_paths3

import (
	"fmt"
	"testing"
)

func TestUniquePaths3(t *testing.T) {
	fmt.Println(uniquePathsIII([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}))
	fmt.Println(uniquePathsIII([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}))
}
