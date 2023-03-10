package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	mysort "gostudy/basic/algorithm/sort"
	datastruct "gostudy/basic/data_struct"
	"sort"
)

func main() {
	a := make([][]int, 0, 1)
	b, _ := json.Marshal(a)
	println(string(b))
	c := make([]int, 0, 1)
	err := json.Unmarshal(b, &c)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", a)
	a = append(a, []int{1})
	fmt.Printf("%+v\n", a)
	list1 := []int{12, 345, 54, 213, 5, 67, 8}
	list2 := []int{12, 345, 54, 213, 5, 67, 8}
	mysort.MergeSort(list1)
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] <= list2[j]
	})
	println(fmt.Sprintf("%v", list1))
	println(fmt.Sprintf("%v", list2))
}

func Test1() {
	type my struct {
		ID int
	}
	m := map[int]my{}
	m[1] = my{
		ID: 1,
	}
	a := m[1]
	a.ID = 2
	print(fmt.Sprintf("%+v", m))
}

func Test2() {
	myHeap := &datastruct.Myheap{8, 2, 10, 4, 5}
	heap.Init(myHeap)
	println(fmt.Sprintf("%+v", myHeap))
	heap.Push(myHeap, 0)
	println(fmt.Sprintf("%+v", myHeap))
	fmt.Printf("%+v", heap.Pop(myHeap))
	println(fmt.Sprintf("%+v", myHeap))
	fmt.Printf("%+v", heap.Pop(myHeap))
	println(fmt.Sprintf("%+v", myHeap))
}
