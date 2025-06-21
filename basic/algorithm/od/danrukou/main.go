package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(strings.TrimSpace("IN 5 10"))
	fmt.Println(strings.Fields(strings.TrimSpace("IN 5 10")))
	reader := bufio.NewReader(os.Stdin)
	numStr, _ := reader.ReadString('\n')
	nums := strings.Split(numStr, " ")
	nums[1] = strings.ReplaceAll(nums[1], "\n", "")
	nums[1] = strings.ReplaceAll(nums[1], "\r", "")
	nums[1] = strings.ReplaceAll(nums[1], " ", "")
	m, _ := strconv.Atoi(nums[0])
	n, _ := strconv.Atoi(nums[1])
	strs := make([][]string, m)
	for i := 0; i < m; i++ {
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		strs[i] = strings.Fields(tmp)
	}
	// 开始处理
	var dfs func(i, j int, count *int, point *[][2]int)
	dfs = func(i, j int, count *int, point *[][2]int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		// 判断是否可以继续
		if strs[i][j] == "X" {
			return
		}
		strs[i][j] = "X"
		if i == 0 || j == 0 || i == m-1 || j == n-1 {
			*point = append(*point, [2]int{i, j})
		}
		*count++
		dfs(i+1, j, count, point)
		dfs(i-1, j, count, point)
		dfs(i, j+1, count, point)
		dfs(i, j-1, count, point)
	}
	res := 0
	tmp := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxO := 0
			point := [][2]int{}
			dfs(i, j, &maxO, &point)
			// 只处理单入口区域（边界点数量为1）
			if maxO == 0 || len(point) != 1 {
				continue
			}
			if maxO < res {
				continue
			} else if maxO == res {
				tmp = append(tmp, point...)
			} else {
				tmp = point
				res = maxO
			}

		}

	}
	if res == 0 {
		fmt.Println("NULL")
		return
	}
	if len(tmp) > 1 {
		fmt.Println(fmt.Sprintf("%d", res))
		return
	}
	fmt.Println(fmt.Sprintf("%d %d %d", tmp[0][0], tmp[0][1], res))
}
