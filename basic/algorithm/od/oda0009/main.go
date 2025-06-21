package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	num, _ := strconv.Atoi(str)
	res := []string{}
	for i := 0; i < num; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strs := strings.Fields(str)
		num1, _ := strconv.Atoi(strs[0])
		num2, _ := strconv.Atoi(strs[1])
		if dfs(num1, num2) {
			res = append(res, "blue")
		} else {
			res = append(res, "red")
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func dfs(n, k int) bool {
	if n == 1 {
		return false
	}
	if k < 1<<(n-2) {
		return !dfs(n-1, k)
	}
	return dfs(n-1, k-1<<(n-2))
}
