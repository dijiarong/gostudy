package count_components

func countComponents(n int, edges [][]int) int {
	uf := NewUF(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	return uf.Count()
}

func NewUF(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  n,
		parent: parent,
	}
}

type UnionFind struct {
	parent []int
	count  int
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

func (un *UnionFind) Count() int {
	return un.count
}

func (un *UnionFind) find(x int) int {
	if x != un.parent[x] {
		un.parent[x] = un.find(un.parent[x])
	}
	return un.parent[x]
}
