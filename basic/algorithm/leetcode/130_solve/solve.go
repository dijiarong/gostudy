package solve

func solve1(board [][]byte) {
	m, n := len(board), len(board[0])
	un := NewUnionFind(m*n + 1)
	root := m * n
	// 处理边缘
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			un.Union(i*n, root)
		}
		if board[i][n-1] == 'O' {
			un.Union(i*n+n-1, root)
		}
	}
	// 处理边缘
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			un.Union(j, root)
		}
		if board[m-1][j] == 'O' {
			un.Union((m-1)*n+j, root)
		}
	}
	// 循环
	nexts := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			for _, v := range nexts {
				x := i + v[0]
				y := j + v[1]
				if board[x][y] == 'O' {
					un.Union(i*n+j, x*n+y)
				}
			}
		}
	}
	// 与root不连接的都需要改成X
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !un.Connect(root, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(count int) *UnionFind {
	parent := make([]int, count)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		count:  count,
	}
}

func (un *UnionFind) Union(a, b int) {
	ra, rb := un.find(a), un.find(b)
	if ra == rb {
		return
	}
	un.parent[ra] = rb
	un.count--
}

func (un *UnionFind) Connect(a, b int) bool {
	return un.find(a) == un.find(b)
}

func (un *UnionFind) find(x int) int {
	if un.parent[x] != x {
		un.parent[x] = un.find(un.parent[x])
	}
	return un.parent[x]
}
