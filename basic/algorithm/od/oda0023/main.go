package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Split(str, "-")
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs = append(strs, strings.Split(str, "-")...)
	none := map[string]int{}
	for i := range strs {
		none[strs[i]]++
	}
	// fmt.Println(none)
	dp := make([][2]int, 12) // 0不包含 1包含
	pai := []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	maxl := 0
	l, r := 0, 0
	for i, v := range pai {
		if i == 0 {
			if none[v] != 4 {
				dp[0][1] = 1
			}
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		if none[v] != 4 {
			dp[i][1] = max(dp[i-1][1]+1, 1)
			if dp[i][1] >= maxl {
				maxl = dp[i][1]
				r = i
				l = r - dp[i][1] + 1
			}
		}
		// fmt.Printf("i:%d, v:%s, dp:%v\n", i, v, dp[i])
	}
	// fmt.Println(dp)
	if r-l < 4 {
		fmt.Println("NO-CHAIN")
		return
	}
	fmt.Println(strings.Join(pai[l:r+1], "-"))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
