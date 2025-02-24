package sort

func QuickSort(list []int) {
	Quic(&list, 0, len(list)-1)
}

func Quic(list *[]int, left, right int) {
	if left >= right {
		return
	}
	// 小于区域
	minArea := left - 1
	maxArea := right
	index := left
	for index < maxArea {
		if (*list)[index] > (*list)[right] {
			Swap(list, index, maxArea-1)
			maxArea -= 1
		} else if (*list)[index] == (*list)[right] {
			index += 1
		} else {
			Swap(list, minArea+1, index)
			minArea += 1
			index += 1
		}
	}
	Swap(list, right, maxArea)
	Quic(list, left, minArea)
	Quic(list, maxArea+1, right)
}
