package sort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	list1 := []int{12, 345, 54, 213, 5, 67, 8}
	MergeSort2(list1)
	print(fmt.Sprintf("%+v", list1))
}
