package datastruct

type Myheap []int

func (m Myheap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m Myheap) Swap(i, j int) {
	tmp := m[j]
	m[j] = m[i]
	m[i] = tmp
}

func (m *Myheap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *Myheap) Pop() any {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func (m Myheap) Len() int {
	return len(m)
}
