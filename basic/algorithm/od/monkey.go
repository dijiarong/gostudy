package od

func eatK(trees []int, h int) int {
	if len(trees) < h {
		return 0
	}
	// 二分法
	total := 0
	for _, tree := range trees {
		total = max(total, tree)
	}
	start := 0
	for start < total {
		mid := (start + total) / 2
		t := getEatTime(trees, mid)
		if t <= h {
			total = mid
		} else {
			start = mid + 1
		}
	}
	if start != 0 {
		return start
	}
	return 0
}

func getEatTime(trees []int, k int) int {
	res := 0
	for _, tree := range trees {
		res += tree / k
		if tree%k != 0 {
			res += 1
		}
	}
	return res
}
