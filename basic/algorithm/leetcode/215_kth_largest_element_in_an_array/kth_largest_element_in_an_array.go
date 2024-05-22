package kthlargestelementinanarray

import "container/heap"

// https://leetcode.cn/problems/kth-largest-element-in-an-array
/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题
*/
func findKthLargest(nums []int, k int) int {
	myh := Myheap(nums)
	heap.Init(&myh)
	for i := 0; i < k; i++ {
		tmp := heap.Pop(&myh)
		if i == k-1 {
			return tmp.(int)
		}
	}
	return myh[k]
}

type Myheap []int

func (m Myheap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m Myheap) Swap(i, j int) {
	tmp := m[j]
	m[j] = m[i]
	m[i] = tmp
}

func (m *Myheap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *Myheap) Pop() any {
	if m.Len() == 0 {
		return nil
	}
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func (m Myheap) Len() int {
	return len(m)
}
