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
	str = strings.TrimSpace(str)
	strs := strings.Fields(str)
	n := len(strs)
	res := make([]string, 0, len(strs))
	var f func(i int)
	f = func(i int) {
		if i > n/2-1 {
			return
		}
		f(2*i + 1)
		f(2*i + 2)
		res = append(res, fmt.Sprintf("%d", i+1))
	}
	f(0)
	fmt.Print(strings.Join(res, " "))
}
