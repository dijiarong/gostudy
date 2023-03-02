package sort

// 冒泡排序
func BubbleSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}
	for i := 1; i <= length; i++ {
		for j := 0; j < length-i; j++ {
			if list[j] > list[j+1] {
				Swap(&list, j, j+1)
			}
		}
	}
	return list
}
