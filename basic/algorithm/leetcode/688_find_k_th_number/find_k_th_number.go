package find_k_th_number

func findKthNumber(m int, n int, k int) int {
	left, right := 1, m*n
	for left < right {
		cnt := 0
		mid := left + (right-left)/2
		for i := 1; i <= m; i++ {
			cnt += min(mid/i, n)
		}
		if cnt < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
