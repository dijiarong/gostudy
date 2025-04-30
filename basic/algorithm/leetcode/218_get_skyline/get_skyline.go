package get_skyline

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i][0] < buildings[j][0]
	})
	h := &tmpHeap{
		arr: [][]int{},
	}
	heap.Init(h)
	for i := range buildings {
		if h.Len() == 0 {
			heap.Push(h, buildings[i])
		}
	}
	return nil
}

type tmpHeap struct {
	arr [][]int
}

func (h *tmpHeap) Push(x any) {
	h.arr = append(h.arr, x.([]int))
}

func (h *tmpHeap) Pop() any {
	tmp := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	return tmp
}

func (h *tmpHeap) Peek() any {
	return h.arr[len(h.arr)-1]
}

func (h *tmpHeap) Less(i, j int) bool {
	return false
}

func (h *tmpHeap) Len() int {
	return len(h.arr)
}

func (h *tmpHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
