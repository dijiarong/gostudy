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
	strs := strings.Split(str, ",")
	works := make([]int, len(strs))
	r := 0
	for i := range strs {
		num, _ := strconv.Atoi(strs[i])
		works[i] = num
		r += num
	}
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	people, _ := strconv.Atoi(str)
	if people == 1 {
		fmt.Print(r)
		return
	}

	// 按工作量从大到小排序
	sort.Slice(works, func(i, j int) bool {
		return works[i] > works[j]
	})

	// 二分搜索的左边界：最大单个任务的工作量
	l := works[0]
	for l < r {
		mid := (l + r) / 2
		if canFinish(works, people, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
}

func canFinish(works []int, people, maxTime int) bool {
	// 使用DFS尝试所有可能的分配方案
	peopleWorkload := make([]int, people)
	return dfs(works, peopleWorkload, 0, maxTime)
}

func dfs(works []int, peopleWorkload []int, taskIdx, maxTime int) bool {
	// 如果所有任务都已分配完成，返回true
	if taskIdx == len(works) {
		return true
	}

	currentWork := works[taskIdx]

	// 尝试将当前任务分配给每个人
	for i := 0; i < len(peopleWorkload); i++ {
		// 如果当前人能承担这个任务
		if peopleWorkload[i]+currentWork <= maxTime {
			peopleWorkload[i] += currentWork

			// 递归尝试分配下一个任务
			if dfs(works, peopleWorkload, taskIdx+1, maxTime) {
				return true
			}

			// 回溯
			peopleWorkload[i] -= currentWork
		}

		// 剪枝：如果当前人的工作量为0，那么后面工作量为0的人也不用尝试了
		// 因为它们是等价的
		if peopleWorkload[i] == 0 {
			break
		}
	}

	return false
}
