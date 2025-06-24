package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str1 := strings.TrimSpace(str)
	str, _ = reader.ReadString('\n')
	str2 := strings.TrimSpace(str)
	r := 0 // c的位置
	l := 0 // a的位置
	for r < len(str2) && l < len(str1) {
		if str2[r] == str1[l] {
			l++
			r++
		} else {
			r++
		}
	}
	if l == len(str1) {
		fmt.Println(r - 1)
		return
	}
	fmt.Println(-1)
}
