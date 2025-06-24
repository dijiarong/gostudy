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
	nums := make([]int, len(strs))
	semap := map[int]int{}
	for i, v := range strs {
		nums[i], _ = strconv.Atoi(v)
	}
	for i := 0; i < len(nums)/2; i++ {
		semap[nums[i*2]] = nums[i*2+1]
	}
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs = strings.Fields(str)
	ql, _ := strconv.Atoi(strs[0])
	wc, _ := strconv.Atoi(strs[1])
	workingmap := map[int]int{}
	q := []int{}
	res1 := 0
	res2 := 0
	for i := 1; i <= 10000; i++ {
		// 查看当前是否有任务结束
		if tmp, ok := workingmap[i]; ok {
			delete(workingmap, i)
			wc += tmp
			if i > res1 {
				res1 = i
			}
		}
		// 无论是否有新任务，都要检查队列中的等待任务
		for len(q) > 0 && wc > 0 {
			workingmap[semap[q[0]]+i]++
			wc--	
			q = q[1:]
		}

		// 然后处理新任务
		if _, ok := semap[i]; ok {
			q = append(q, i)
			// 再次检查是否可以立即执行
			for len(q) > 0 && wc > 0 {
				workingmap[i+semap[q[0]]]++
				wc--
				q = q[1:]
			}
			// 队列满了就丢弃
			if len(q) > ql {
				q = q[1:]
				res2++
			}
		}
	}
	fmt.Printf("%d %d", res1, res2)
}
