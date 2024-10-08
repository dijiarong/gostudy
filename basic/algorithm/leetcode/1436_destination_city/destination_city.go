package destinationcity

func destCity(paths [][]string) string {
	// 记录路径
	pathMap := make(map[string]bool)
	for _, v := range paths {
		pathMap[v[0]] = true
	}

	for _, v := range paths {
		if !pathMap[v[1]] {
			return v[1]
		}
	}
	return ""
}
