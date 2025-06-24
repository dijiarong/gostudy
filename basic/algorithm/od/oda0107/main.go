package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	maxTime, _ := strconv.Atoi(str)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Fields(str)
	sliders := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		sliders[i], _ = strconv.Atoi(strs[i])
	}
	sort.Ints(sliders)
	dp := make([]int, len(strs))
	dp[0] = sliders[0]
	dp[1] = getMax(sliders[0], sliders[1])
	for i := 2; i < len(sliders); i++ {
		dp[i] = min(dp[i-1]+sliders[0]+getMax(sliders[0], sliders[i]), dp[i-2]+sliders[0]+getMax(sliders[i-1], sliders[i])+sliders[1]+getMax(sliders[0], sliders[1]))
		if dp[i] > maxTime {
			fmt.Printf("%d %d", i, dp[i-1])
			return
		}
	}
	fmt.Printf("%d %d", len(dp), dp[len(dp)-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMax(a, b int) int {
	if a*10 < b {
		return a * 10
	}
	return b
}
