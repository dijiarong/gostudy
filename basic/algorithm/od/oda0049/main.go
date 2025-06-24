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
	strs := strings.Fields(str)
	n, _ := strconv.Atoi(strs[0])
	m, _ := strconv.Atoi(strs[1])
	nums := make([][]string, n)
	uesd := make([][]bool, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strs = strings.Fields(str)
		nums[i] = strs
		uesd[i] = make([]bool, m)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if uesd[i][j] || nums[i][j] == "2" {
			return
		}
		nums[i][j] = "1"
		uesd[i][j] = true
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	nums[0][0] = "1"
	dfs(0, 0)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums[i][j] != "1" {
				res++
			}
		}
	}
	fmt.Println(res)
}
