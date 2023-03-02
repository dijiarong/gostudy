package sort

type sort struct{}

var Sort sort

func Swap(list *[]int, a int, b int) {
	t := (*list)[a]
	(*list)[a] = (*list)[b]
	(*list)[b] = t
}
