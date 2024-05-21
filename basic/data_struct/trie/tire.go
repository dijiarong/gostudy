package trie

type Trie struct {
	children [26]*Trie
	isEnd    bool
	passCnt  int
}

func New() *Trie {
	return &Trie{}
}

// 插入word
func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}
	node := t
	for _, c := range word {
		tmp := c - 'a'
		if node.children[tmp] == nil {
			node.children[tmp] = &Trie{}
		}
		node.children[tmp].passCnt++
		node = node.children[tmp]
	}
	node.isEnd = true
}

// 查询word是否存在
func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.isEnd
}

// 查询word的节点
func (t *Trie) search(word string) *Trie {
	node := t
	for _, c := range word {
		tmp := c - 'a'
		if node.children[tmp] == nil {
			return nil
		}
		node = node.children[tmp]
	}
	return node
}

// 查询是否存在以prefix为前缀的单词
func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}

// 查询以prefix为前缀的单词数量
func (t *Trie) PassCount(prefix string) int {
	node := t.search(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

// 删除word
func (t *Trie) Delete(word string) {
	if !t.Search(word) {
		return
	}
	node := t
	for _, c := range word {
		tmp := c - 'a'
		node.children[tmp].passCnt--
		if node.children[tmp].passCnt == 0 {
			node.children[tmp] = nil
			return
		}
		node = node.children[tmp]
	}
	node.isEnd = false
}
