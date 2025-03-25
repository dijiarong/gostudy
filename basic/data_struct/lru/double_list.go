package lru

type node struct {
	key, val  int
	pre, next *node
}

type doubleList struct {
	head, tail *node
	size       int
}

func NewDoubleList() *doubleList {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.pre = head
	return &doubleList{
		head, tail, 0,
	}
}

func (l *doubleList) Remove(node *node) *node {
	node.pre.next = node.next
	node.next.pre = node.pre
	l.size--
	return node
}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (l *doubleList) RemoveFirst() *node {
	if l.head.next == l.tail {
		return nil
	}
	first := l.head.next
	l.Remove(first)
	return first
}

func (l *doubleList) AddLast(node *node) {
	l.tail.pre.next = node
	node.pre = l.tail.pre
	node.next = l.tail
	l.tail.pre = node
	l.size++
}

func (l *doubleList) AddFirst(node *node) {
	l.tail.pre.next = node
	node.pre = l.tail.pre
	node.next = l.tail
	l.tail.pre = node
	l.size++
}

func (l *doubleList) Size() int {
	return l.size
}
