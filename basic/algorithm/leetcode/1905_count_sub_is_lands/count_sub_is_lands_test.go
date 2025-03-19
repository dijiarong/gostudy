package count_sub_is_lands

import (
	"fmt"
	"testing"
)

func TestCountSubIslands(t *testing.T) {
	fmt.Print(countSubIslands([][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1},
	}, [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
	}))
}
