package find_words

type node struct {
	childs [26]*node
	word   string
}

type trie struct {
	root *node
}

func (t *trie) insert(word string) {
	n := t.root
	for _, v := range word {
		if n.childs[v-'a'] == nil {
			n.childs[v-'a'] = &node{
				childs: [26]*node{},
			}
		}
		n = n.childs[v-'a']
	}
	n.word = word
}

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	t := trie{
		root: &node{
			childs: [26]*node{},
		},
	}
	for _, v := range words {
		t.insert(v)
	}
	nexts := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	res := []string{}
	var backTrack func(i, j int, nd *node)
	backTrack = func(i, j int, nd *node) {
		if i < 0 || j < 0 || i >= m || j >= n || nd == nil || board[i][j] == '.' {
			return
		}
		original := board[i][j]
		nd = nd.childs[original-'a']
		if nd.word != "" {
			res = append(res, nd.word)
			nd.word = ""
		}
		board[i][j] = '.'
		for _, v := range nexts {
			x, y := i+v[0], j+v[1]
			backTrack(x, y, nd)
		}
		board[i][j] = original
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backTrack(i, j, t.root)
		}
	}
	return res
}
