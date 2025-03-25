package range_sum_query_mutable

import (
	"fmt"
	"testing"
)

func TestRangeSumQueryMutable(t *testing.T) {
	nums := []int{1, 3, 5}
	array := Constructor(nums)
	fmt.Println(array.SumRange(0, 2))
	array.Update(1, 2)
	fmt.Println(array.SumRange(0, 2))
}
