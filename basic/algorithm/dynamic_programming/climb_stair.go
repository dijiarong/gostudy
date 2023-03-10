package dynamicprogramming

func ClimbStairs(n int) int {
	a, b := 1, 2
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}
	for i := 3; i < n; i++ {
		c := a + b
		a = b
		b = c
	}
	return a + b
}
