package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	Key int
	Val []int
}

type myheap struct {
	arr []*node
}

func (m *myheap) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *myheap) Len() int {
	return len(m.arr)
}

func (m *myheap) Less(i, j int) bool {
	return m.arr[i].Key > m.arr[j].Key
}

func (m *myheap) Push(x any) {
	m.arr = append(m.arr, x.(*node))
}

func (m *myheap) Pop() any {
	old := m.arr
	n := len(old)
	item := old[n-1]
	m.arr = old[0 : n-1]
	return item
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	num, _ := strconv.Atoi(str)
	// 初始化堆dict
	hmap := map[string]*myheap{
		"1": &myheap{
			arr: []*node{},
		},
		"2": &myheap{
			arr: []*node{},
		},
		"3": &myheap{
			arr: []*node{},
		},
		"4": &myheap{
			arr: []*node{},
		},
		"5": &myheap{
			arr: []*node{},
		},
	}
	nodeMap := map[string]map[int]*node{
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
	}
	fileCounter := 0
	for i := 0; i < num; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strs := strings.Fields(str)
		h := hmap[strs[1]]
		if strs[0] == "OUT" {
			pstr := "NULL"
			if h.Len() != 0 {
				// 查看顶点
				pstr = fmt.Sprintf("%d", h.arr[0].Val[0])
				if len(h.arr[0].Val) == 1 {
					// 删除节点记录
					delete(nodeMap[strs[1]], h.arr[0].Key)
					heap.Pop(h)
				} else {
					h.arr[0].Val = h.arr[0].Val[1:]
				}
			}
			fmt.Println(pstr)
		} else {
			fileCounter++
			// 是否在堆中，如果不在初始化进去，在的话直接将index放入val
			num, _ := strconv.Atoi(strs[2])
			tmp, ok := nodeMap[strs[1]][num]
			if ok {
				tmp.Val = append(tmp.Val, fileCounter)
			} else {
				newNode := &node{
					Key: num,
					Val: []int{fileCounter},
				}
				heap.Push(h, newNode)
				nodeMap[strs[1]][num] = newNode
			}
		}
	}
}
