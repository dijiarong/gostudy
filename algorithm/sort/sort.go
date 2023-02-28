package sort

type sort struct{}

var Sort sort

func (sort) Swap(list []int, a int, b int) {
	t := list[a]
	list[a] = list[b]
	list[b] = t
}
