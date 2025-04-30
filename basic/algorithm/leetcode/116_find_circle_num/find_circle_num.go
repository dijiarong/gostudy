package find_circle_num

func findCircleNum(isConnected [][]int) int {
	// 使用dfs算法，将连接的省份递归取消链接
	res := 0
	m := len(isConnected)
	visted := make([]bool, m)
	var dfs func(from int)
	dfs = func(from int) {
		for to, v := range isConnected[from] {
			if v == 1 && !visted[to] {
				dfs(to)
			}
		}
		visted[from] = true
	}
	for from := 0; from < m; from++ {
		if !visted[from] {
			res++
			dfs(from)
		}
	}
	return res
}
