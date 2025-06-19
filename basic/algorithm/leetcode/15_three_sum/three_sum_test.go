package three_sum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
