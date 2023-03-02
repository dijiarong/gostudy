package sort

func MergeSort(list []int) {
	if list == nil || len(list) < 2 {
		return
	}
	Process(&list, 0, len(list)-1)
}

// 非递归版本
func MergeSort2(list []int) {
	// 边界条件
	if list == nil || len(list) < 2 {
		return
	}
	// 以step进行调整，相当于递归
	length := len(list)
	step := 1
	for step < length {
		// 每个步长都应该从0开始
		left := 0
		// 当左边到达最后一位时结束
		for left < length {
			// 定义mid 和右边界
			mid := 0
			right := 0
			if length-left > step {
				mid = left + step - 1
			} else {
				mid = length - 1
			}
			if mid == length-1 {
				break
			}
			if length-mid > step {
				right = mid + step
			} else {
				right = length - 1
			}
			Merge(&list, left, right, mid)
			if right == length-1 {
				break
			}
			left = right + 1
		}
		// step每次乘2
		step *= 2
	}
}

func Process(list *[]int, l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	Process(list, l, mid)
	Process(list, mid+1, r)
	Merge(list, l, r, mid)
}

func Merge(list *[]int, l, r, m int) {
	help := make([]int, r-l+1)
	i, p1, p2 := 0, l, m+1
	for p1 <= m && p2 <= r {
		if (*list)[p1] <= (*list)[p2] {
			help[i] = (*list)[p1]
			p1++
		} else {
			help[i] = (*list)[p2]
			p2++
		}
		i++
	}
	for p1 <= m {
		help[i] = (*list)[p1]
		i++
		p1++
	}
	for p2 <= r {
		help[i] = (*list)[p2]
		i++
		p2++
	}
	for k, v := range help {
		(*list)[k+l] = v
	}
}
