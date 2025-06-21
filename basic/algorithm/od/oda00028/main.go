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
	nums := make([]uint32, len(strs))
	for i, v := range strs {
		tmp, _ := strconv.Atoi(v)
		nums[i] = uint32(tmp)
	}
	c, b := nums[0], nums[1]
	cnt := make([]int, c)
	for i := 2; i < len(nums); i++ {
		tmp := uint32(byte(nums[i])) + uint32(byte(nums[i]>>8)) + uint32(byte(nums[i]>>16)) + uint32(byte(nums[i]>>24))
		if tmp%b < c {
			cnt[tmp%b]++
		}
	}
	res := 0
	for _, v := range cnt {
		if v > res {
			res = v
		}
	}
	fmt.Println(fmt.Sprintf("%d", res))
}
