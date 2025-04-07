package candy

func candy(ratings []int) int {
	n := len(ratings)
	need := make([]int, n)
	for i := range need {
		need[i] = 1
	}
	for i := 0; i < n-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			need[i-1] = need[i] + 1
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += need[i]
	}
	return res
}
