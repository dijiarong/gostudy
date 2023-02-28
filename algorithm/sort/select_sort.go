package sort

// 选择排序:记录最小或者最大那个索引，每一次循环都可以找到一个
func SelectSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}
	for i := 0; i < length; i++ {
		maxIndex := i
		for j := i + 1; j < length; j++ {
			if list[j] < list[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			Sort.Swap(list, i, maxIndex)
		}
	}
	return list
}
