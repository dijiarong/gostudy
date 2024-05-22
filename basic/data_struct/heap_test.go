package datastruct

import (
	"container/heap"
	"errors"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Myheap{}
	heap.Init(h)
	nums := []int{3, 2, 1, 5, 6, 4}
	nums1 := []int{3, 2, 1, 5, 6, 4}
	for _, v := range nums {
		heap.Push(h, v)
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] >= nums1[j]
	})
	for _, v := range nums1 {
		if v != heap.Pop(h).(int) {
			panic(errors.New(""))
		}
	}
}
