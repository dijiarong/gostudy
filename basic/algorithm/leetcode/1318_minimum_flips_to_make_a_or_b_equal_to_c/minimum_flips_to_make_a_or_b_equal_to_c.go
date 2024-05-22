package minimumflipstomakeaorbequaltoc

// https://leetcode.cn/problems/minimum-flips-to-make-a-or-b-equal-to-c
/*
给你三个正整数 a、b 和 c。

你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。

「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。
*/
func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 31; i++ {
		tmpa, tmpb, tmpc := a&1, b&1, c&1
		if tmpc == 0 {
			res = res + tmpa + tmpb
		} else if tmpa+tmpb < 1 {
			res += 1
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return res
}
