package range_sum_query_immutable

import (
	"fmt"
	"testing"
)

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2)) // return 1 because (3 + (-5)) = 1
	fmt.Println(obj.SumRange(2, 5)) // return -1 because ((3) + (-5) + (2) + (-1)) = -1
	fmt.Println(obj.SumRange(0, 5)) // return -3
}
