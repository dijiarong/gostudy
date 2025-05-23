package skip_list

import "math/rand"

type SkipList struct {
	head *node
}

type node struct {
	Key   int
	Value int
	Nexts []*node
}

func (s *SkipList) Get(key int) (int, bool) {
	if n := s.search(key); n != nil {
		return n.Value, true
	}
	return -1, false
}

func (s *SkipList) Put(key, value int) {
	if n := s.search(key); n != nil {
		n.Value = value
		return
	}
	// 如果不存在需要插入
	// 1.确认层数
	level := s.roll()
	// 2.如果该节点层数超过了已有层数需要处理一下
	for len(s.head.Nexts)-1 < level {
		s.head.Nexts = append(s.head.Nexts, nil)
	}
	// 3.初始化新节点
	newNode := &node{
		Key:   key,
		Value: value,
		Nexts: make([]*node, level+1),
	}
	n := s.head
	// 4.一层一层处理新整节点
	for level := level; level >= 0; level-- {
		// 找到最接近的节点
		for n.Nexts[level] != nil && n.Nexts[level].Key < key {
			n = n.Nexts[level]
		}
		// 插入该节点
		newNode.Nexts[level] = n.Nexts[level]
		n.Nexts[level] = newNode
	}
}

func (s *SkipList) Del(key int) {
	if n := s.search(key); n == nil {
		return
	}
	n := s.head
	level := len(s.head.Nexts) - 1
	for level := level; level >= 0; level-- {
		for n.Nexts[level] != nil && n.Nexts[level].Key < key {
			n = n.Nexts[level]
		}
		if n.Nexts[level] == nil || n.Nexts[level].Key > key {
			continue
		}
		n.Nexts[level] = n.Nexts[level].Nexts[level]
	}
	incre := 0
	for level := len(s.head.Nexts) - 1; level > 0 && s.head.Nexts[level] == nil; level-- {
		incre++
	}
	if incre > 0 {
		s.head.Nexts = s.head.Nexts[:len(s.head.Nexts)-incre]
	}
}

func (s *SkipList) Range(start, end int) [][2]int {
	res := [][2]int{}
	// 找到开始节点
	n := s.celling(start)
	if n == nil {
		return res
	}
	for cur := n; cur != nil && cur.Key <= end; cur = cur.Nexts[0] {
		res = append(res, [2]int{cur.Key, cur.Value})
	}
	return res
}

func (s *SkipList) Celling(key int) [2]int {
	res := [2]int{}
	n := s.celling(key)
	if n != nil {
		res[0] = n.Key
		res[1] = n.Value
	}
	return res
}

// 大于key的最小值
func (s *SkipList) celling(key int) *node {
	level := len(s.head.Nexts) - 1
	n := s.head
	// 一层一层找
	for level := level; level >= 0; level-- {
		// 沿着当前层找到最接近key的node
		for n.Nexts[level] != nil && n.Nexts[level].Key < key {
			n = n.Nexts[level]
		}
		// 判断一下，如果相等直接返回
		if n.Nexts[level] != nil && n.Nexts[level].Key == key {
			return n.Nexts[level]
		}
	}
	return n.Nexts[0]
}

func (s *SkipList) Floor(key int) [2]int {
	res := [2]int{}
	n := s.floor(key)
	if n != nil {
		res[0] = n.Key
		res[1] = n.Value
	}
	return res
}

// 小于key的最大值
func (s *SkipList) floor(key int) *node {
	level := len(s.head.Nexts) - 1
	n := s.head
	// 一层一层找
	for level := level; level >= 0; level-- {
		// 沿着当前层找到最接近key的node
		for n.Nexts[level] != nil && n.Nexts[level].Key < key {
			n = n.Nexts[level]
		}
		// 判断一下，如果相等直接返回
		if n.Nexts[level] != nil && n.Nexts[level].Key == key {
			return n.Nexts[level]
		}
	}
	return n
}

// 限制最大层数，避免层数过高
const MaxLevel = 16

func (s *SkipList) roll() int {
	var level int
	// 每次投出 0，则层数加 1，但要限制最大层数
	for rand.Intn(2) == 0 && level < MaxLevel {
		level++
	}
	return level
}

func (s *SkipList) search(key int) *node {
	level := len(s.head.Nexts) - 1
	n := s.head
	// 一层一层找
	for level := level; level >= 0; level-- {
		// 沿着当前层找到最接近key的node
		for n.Nexts[level] != nil && n.Nexts[level].Key < key {
			n = n.Nexts[level]
		}
		// 判断一下，如果相等直接返回
		if n.Nexts[level] != nil && n.Nexts[level].Key == key {
			return n.Nexts[level]
		}
	}
	return nil
}
