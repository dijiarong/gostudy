package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Fields(str)
	// 滑动窗口
	n, _ := strconv.Atoi(strs[0]) // Hang
	m, _ := strconv.Atoi(strs[1]) // Lie
	nums := make([][]string, n)
	for i := 0; i < n; i++ {
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strs = strings.Fields(str)
		nums[i] = strs
	}
	// fmt.Println(nums)
	str, _ = reader.ReadString('\n')
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs = strings.Fields(str)
	need := map[string]int{}
	for _, v := range strs {
		need[v]++
	}
	// fmt.Println(need)
	windows := map[string]int{}
	// 从头开始
	res := math.MaxInt
	l := 0
	r := 0
	for r < m {
		for j := 0; j < n; j++ {
			_, ok := need[nums[j][r]]
			if !ok {
				continue
			}
			windows[nums[j][r]]++
		}
		r++
		// fmt.Print("windows:扩展")
		// fmt.Println(windows)
		for check(need, windows) {
			if r-l < res {
				res = r - l
			}
			for j := 0; j < n; j++ {
				_, ok := need[nums[j][l]]
				if !ok {
					continue
				}
				windows[nums[j][l]]--
			}
			l++
		}
	}
	if res == math.MaxInt {
		fmt.Print(-1)
	} else {
		fmt.Println(res)
	}
}

func check(target, cur map[string]int) bool {
	if len(target) != len(cur) {
		return false
	}
	for k, v := range target {
		if v > cur[k] {
			return false
		}
	}
	return true
}
