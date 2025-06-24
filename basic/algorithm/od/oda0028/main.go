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
	n, _ := strconv.Atoi(str)
	m := make([][]string, n)
	si, sj := 0, 0
	candyMap := make([][]int, n)
	for i := 0; i < n; i++ {
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		m[i] = strings.Fields(str)
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == "-3" {
				si, sj = i, j
			}
		}
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			tmp[i] = -1
		}
		candyMap[i] = tmp
	}
	nexts := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := 0
	q := [][3]int{{si, sj, 0}}
	candyMap[si][sj] = 0
	find := false
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			// 判断下一个节点是否已经使用过且不越界
			for _, next := range nexts {
				i, j := cur[0]+next[0], cur[1]+next[1]
				if i < 0 || j < 0 || i >= n || j >= n {
					continue
				}
				if m[i][j] == "-1" {
					continue
				}
				if m[i][j] == "-2" {
					find = true
					res = max(res, cur[2])
					continue
				}
				newCandy := cur[2]
				tmp, _ := strconv.Atoi(m[i][j])
				newCandy += tmp
				if candyMap[i][j] == -1 || newCandy > candyMap[i][j] {
					q = append(q, [3]int{i, j, cur[2] + tmp})
					candyMap[i][j] = newCandy
				}
			}
			q = q[1:]
		}
		if find {
			fmt.Println(res)
			return
		}
	}
	fmt.Println(-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
