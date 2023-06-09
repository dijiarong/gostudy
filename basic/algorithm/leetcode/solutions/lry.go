package solutions

import "container/list"

type Object struct {
	key, val int
}

type LRUCache struct {
	len, cap   int
	objectMap  map[int]*list.Element
	objectList *list.List
}

func Constructer(cap int) *LRUCache {
	return &LRUCache{
		len:        0,
		cap:        cap,
		objectMap:  make(map[int]*list.Element),
		objectList: list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if e, ok := lru.objectMap[key]; ok {
		// 移动查询元素位置
		lru.objectList.MoveToFront(e)
		return e.Value.(*Object).val
	}
	return -1
}

func (lru *LRUCache) Put(key, val int) {
	e, ok := lru.objectMap[key]
	if ok {
		lru.objectList.MoveToFront(e)
		o := e.Value.(*Object)
		o.val = val
		return
	}
	// o := &Object{
	// 	key: key,
	// 	val: val,
	// }
	if lru.len == lru.cap {

	}
}
