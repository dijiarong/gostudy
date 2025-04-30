package skip_list

import (
	"math/rand"
)

type SkipList struct {
	head *node
}

type node struct {
	Key   int
	Value int
	Nexts []*node
}

func (s *SkipList) Get(key int) (int, bool) {
	if _node := s.search(key); _node != nil {
		return _node.Value, true
	}
	return 0, false
}

func (s *SkipList) search(key int) *node {
	move := s.head
	// 从高层一层一层读取
	level := len(s.head.Nexts) - 1
	for level := level; level >= 0; level-- {
		// 同层查询
		for move.Nexts[level] != nil && move.Nexts[level].Key < key {
			move = move.Nexts[level]
		}
		// 上面循环结束，move同层下一个元素的key如果==key
		if move.Nexts[level] != nil && move.Nexts[level].Key == key {
			return move.Nexts[level]
		}
	}
	return nil
}

func (s *SkipList) Put(key int, value int) {
	// 查询如果存在则更新
	if _node := s.search(key); _node != nil {
		_node.Value = value
		return
	}
	// 随机层数
	level := s.roll()
	move := s.head
	newNode := &node{
		Key:   key,
		Value: value,
		Nexts: make([]*node, level+1),
	}
	for level := level; level >= 0; level-- {
		// 同层查询
		for move.Nexts[level] != nil && move.Nexts[level].Key < key {
			move = move.Nexts[level]
		}
		newNode.Nexts[level] = move.Nexts[level]
		move.Nexts[level] = newNode
	}
}

func (s *SkipList) Del(key int) {
	// 查询如果存在则更新
	if _node := s.search(key); _node == nil {
		return
	}
	move := s.head
	level := len(move.Nexts) - 1
	for level := level; level >= 0; level-- {
		for move.Nexts[level] != nil && move.Nexts[level].Key < key {
			move = move.Nexts[level]
		}
		if move.Nexts[level] == nil || move.Nexts[level].Key > key {
			continue
		}
		move.Nexts[level] = move.Nexts[level].Nexts[level]
	}
	// 删除元素后要检查是否要降低高度
	incre := 0
	for level := len(s.head.Nexts); level > 0 && s.head.Nexts[level] == nil; level-- {
		incre++
	}
	s.head.Nexts = s.head.Nexts[:len(s.head.Nexts)-incre]
}

func (s *SkipList) roll() int {
	var level int
	// 每次投出 1，则层数加 1
	for rand.Intn(2) == 0 {
		level++
	}
	return level
}

func (s *SkipList) Range(start, end int) [][2]int {
	res := [][2]int{}
	// 首先获取离start最近的
	startnode := s.celling(start)
	if startnode == nil {
		return res
	}
	for move := startnode; move != nil && move.Key <= end; move = move.Nexts[0] {
		res = append(res, [2]int{move.Key, move.Value})
	}
	return res
}

func (s *SkipList) Celling(stat int) [2]int {
	res := [2]int{}
	n := s.celling(stat)
	if n == nil {
		return res
	}
	// 首先获取离start最近的
	res[0] = n.Key
	res[1] = n.Value
	return res
}

func (s *SkipList) celling(stat int) *node {
	move := s.head
	for level := len(s.head.Nexts) - 1; level >= 0; level-- {
		for move.Nexts[level] != nil && move.Nexts[level].Key < stat {
			move = move.Nexts[level]
		}
		if move.Nexts[level] != nil && move.Nexts[level].Key == stat {
			return move.Nexts[level]
		}
	}
	// 首先获取离start最近的
	return move.Nexts[0]
}

func (s *SkipList) Floor(key int) [2]int {
	res := [2]int{}
	n := s.floor(key)
	if n == nil {
		return res
	}
	res[0] = n.Key
	res[1] = n.Value
	return res
}

func (s *SkipList) floor(stat int) *node {
	move := s.head
	for level := len(s.head.Nexts) - 1; level >= 0; level-- {
		for move.Nexts[level] != nil && move.Nexts[level].Key < stat {
			move = move.Nexts[level]
		}
		if move.Nexts[level] != nil && move.Nexts[level].Key == stat {
			return move.Nexts[level]
		}
	}
	// 首先获取离start最近的
	return move
}
