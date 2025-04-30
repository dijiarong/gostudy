package row_with_maximum_ones

func rowAndMaximumOnes(mat [][]int) []int {
	zeroNumMax := 0
	index := -1
	for i, v := range mat {
		curZeroNum := 0
		for _, j := range v {
			if j == 1 {
				curZeroNum++
			}
		}
		if curZeroNum > zeroNumMax {
			zeroNumMax = curZeroNum
			index = i
		}
	}
	return []int{index, zeroNumMax}
}
