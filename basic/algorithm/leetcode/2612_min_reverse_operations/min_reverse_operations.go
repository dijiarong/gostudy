package min_reverse_operations

func minReverseOperations(n int, p int, banned []int, k int) []int {
	res := make([]int, n)
	// 初始化结果数组，默认值为 -1
	for i := range res {
		res[i] = -1
	}
	bannedMap := map[int]bool{}
	for _, v := range banned {
		bannedMap[v] = true
	}
	res[p] = 0
	q := []int{p}
	step := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			for start := max(0, cur-k+1); start <= min(cur, n-k); start++ {
				curMove := start + start + k - 1 - cur
				// 检查新位置是否合法且未被访问
				if !bannedMap[curMove] && res[curMove] == -1 {
					q = append(q, curMove)
					bannedMap[curMove] = true
					// 更新结果数组
					res[curMove] = step
				}
			}
		}
		step++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
