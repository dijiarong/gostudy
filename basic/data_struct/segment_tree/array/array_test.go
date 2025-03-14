package segment_tree

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	// 示例，创建一棵求和线段树
	st := NewArraySegmentTree(arr, func(a, b int) int { return a + b })

	result := st.Query(1, 3)
	fmt.Println(result) // 3 + 5 + 7 = 15
	st.Update(2, 10)
	result = st.Query(1, 3)
	fmt.Println(result) // 3 + 10 + 7 = 20
}
