package segment_tree

import (
	"fmt"
	"testing"
)

func TestAssignSegmentTree(t *testing.T) {
	// 构建区间 [0, 9] 的赋值线段树，初始值均为 0
	tree := NewAssignSegmentTree(0, 9, 0)
	// [0,0,0,0,0,0,0,0,0,0]

	tree.RangeUpdate(2, 5, 7)
	// [0,0,7,7,7,7,7,0,0,0]

	// 7 * 4 = 28
	fmt.Println("sum of [0,9]:", tree.Query(0, 9), "expected: 28")

	tree.RangeUpdate(0, 3, 5)
	// [5,5,5,5,7,7,0,0,0,0]

	// 5 * 4 + 7 * 2 = 34
	fmt.Println("sum of [0,5]:", tree.Query(0, 5), "expected: 34")

	tree.RangeUpdate(5, 9, 2)
	// [5,5,5,5,7,2,2,2,2,2]

	// 5 * 4 + 7 * 2 + 2 * 5 = 37
	fmt.Println("sum of [0,9]:", tree.Query(0, 9), "expected: 37")

	// 单点查询索引 5 的值
	fmt.Println("value of index 5:", tree.Query(5, 5), "expected: 2")

	// 5 + 5 = 10
	fmt.Println("sum of [1,2]:", tree.Query(1, 2), "expected: 10")
}
