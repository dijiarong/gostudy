package open_lock

func openLock(deadends []string, target string) int {
	step := 0
	start := "0000"
	visited := map[string]bool{}
	for _, v := range deadends {
		visited[v] = true
	}
	if visited[start] {
		return -1
	}
	q := []string{start}
	visited[start] = true
	for len(q) > 0 {
		tmp := len(q)
		for i := 0; i < tmp; i++ {
			cur := q[0]
			if cur == target {
				return step
			}
			q = q[1:]
			for _, v := range getNexts(cur) {
				if !visited[v] {
					q = append(q, v)
					visited[v] = true
				}
			}
		}
		step++
	}
	return -1
}

func upIndex(s string, index int) string {
	res := []rune(s)
	if res[index] == '9' {
		res[index] = '0'
	} else {
		res[index]++
	}
	return string(res)
}

func downIndex(s string, index int) string {
	res := []byte(s)
	if res[index] == '0' {
		res[index] = '9'
	} else {
		res[index]--
	}
	return string(res)
}

func getNexts(s string) []string {
	res := []string{}
	for i := 0; i < 4; i++ {
		res = append(res, upIndex(s, i))
		res = append(res, downIndex(s, i))
	}
	return res
}
