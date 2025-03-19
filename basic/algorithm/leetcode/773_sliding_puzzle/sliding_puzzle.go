package sliding_puzzle

import (
	"fmt"
	"strings"
)

func slidingPuzzle(board [][]int) int {
	res := 0
	start, target := "", "123450"
	for _, v := range board {
		for _, w := range v {
			start += fmt.Sprintf("%d", w)
		}
	}
	visited := map[string]bool{
		start: true,
	}
	q := []string{start}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			tmp := q[0]
			if tmp == target {
				return res
			}
			q = q[1:]
			for _, next := range getNexts(tmp) {
				if !visited[next] {
					visited[next] = true
					q = append(q, next)
				}
			}
		}
		res++
	}
	return res
}

var nextMap = [][]int{
	{1, 3},
	{0, 2, 4},
	{1, 5},
	{0, 4},
	{1, 3, 5},
	{2, 4},
}

func getNexts(s string) []string {
	res := []string{}
	index := strings.Index(s, "0")
	for _, v := range nextMap[index] {
		res = append(res, swap(s, index, v))
	}
	return res
}

func swap(s string, i, j int) string {
	res := []rune(s)
	res[i], res[j] = res[j], res[i]
	return string(res)
}
