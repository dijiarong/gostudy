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
	num, _ := strconv.Atoi(str)
	points := [][2]int{}
	for i := 0; i < num; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strs := strings.Split(str, " ")
		start, _ := strconv.Atoi(strs[0])
		end, _ := strconv.Atoi(strs[1])
		points = append(points, [2]int{start, end})
	}
	doubles := [][2]int{}
	// 顺便去重
	m := map[int]map[int]bool{}
	for i := 0; i < len(points); i++ {
		s1 := points[i]
		for j := i + 1; j < len(points); j++ {
			s2 := points[j]
			strat := max(s1[0], s2[0])
			end := min(s1[1], s2[1])
			if strat <= end {
				if _, ok := m[strat]; !ok {
					m[strat] = map[int]bool{}
				}
				if !m[strat][end] {
					doubles = append(doubles, [2]int{strat, end})
					m[strat][end] = true
				}
			}
		}
	}
	if len(doubles) == 0 {
		fmt.Print("None")
		return
	}
	// 求出重叠区间后，再次处理
	sort.Slice(doubles, func(i, j int) bool {
		if doubles[i][0] == doubles[j][0] {
			return doubles[i][1] < doubles[j][1]
		}
		return doubles[i][0] < doubles[j][0]
	})
	// 计算重叠
	cur := doubles[0]
	res := [][2]int{}
	for i := 1; i < len(doubles); i++ {
		next := doubles[i]
		if cur[1] >= next[0] {
			cur[1] = max(cur[1], next[1])
		} else {
			res = append(res, cur)
			cur = next
		}
	}
	res = append(res, cur)
	for _, v := range res {
		fmt.Println(fmt.Sprintf("%d %d", v[0], v[1]))
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
