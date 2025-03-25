package lru

type LRU struct {
	cache *doubleList
	_map  map[int]*node
	cap   int
}

func NewLRUCache(cap int) *LRU {
	return &LRU{
		cache: NewDoubleList(),
		_map:  make(map[int]*node, cap),
		cap:   cap,
	}
}

// 定义三个基础方法
func (this *LRU) makeRecently(key int) {
	if this.cache.size == 0 {
		return
	}
	tmp, ok := this._map[key]
	if !ok {
		return
	}
	this.cache.Remove(tmp)
	this.cache.AddLast(tmp)
}

func (this *LRU) addRecently(key, val int) {
	x := &node{
		key: key,
		val: val,
	}
	// 链表尾部就是最近使用的元素
	this.cache.AddLast(x)
	// 别忘了在 map 中添加 key 的映射
	this._map[key] = x
}

func (this *LRU) deleteKey(key int) {
	if this.cache.size == 0 {
		return
	}
	tmp, ok := this._map[key]
	if !ok {
		return
	}
	this.cache.Remove(tmp)
	delete(this._map, key)
}

// 移除末尾元素
func (this *LRU) removeLast() {
	deletedNode := this.cache.RemoveFirst()
	// 同时别忘了从 map 中删除它的 key
	deletedKey := deletedNode.key
	delete(this._map, deletedKey)
}
