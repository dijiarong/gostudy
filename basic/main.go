package main

import (
	"container/heap"
	"fmt"
	datastruct "myself/data_struct"
)

func main() {
	myHeap := &datastruct.Myheap{1, 2, 3, 4, 5}
	heap.Init(myHeap)
	println(fmt.Sprintf("%+v", myHeap))
	heap.Push(myHeap, 0)
	println(fmt.Sprintf("%+v", myHeap))
	println(heap.Pop(myHeap))
	println(fmt.Sprintf("%+v", myHeap))
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
