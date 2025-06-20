package generate_parentheses

func generateParenthesis(n int) []string {
	res := []string{}
	var track string
	var trackback func(left, right int)
	trackback = func(left, right int) {
		if left < right {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			res = append(res, track)
			return
		}
		track += "("
		trackback(left-1, right)
		track = track[:len(track)-1]
		track += ")"
		trackback(left, right-1)
		track = track[:len(track)-1]
	}
	trackback(n, n)
	return res
}
