package destinationcity

import (
	"fmt"
	"testing"
)

func TestDestCity(t *testing.T) {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	// 记录路径
	pathMap := make(map[string]bool)
	for _, v := range paths {
		pathMap[v[0]] = true
	}

	for _, v := range paths {
		if !pathMap[v[1]] {
			fmt.Println(v[1])
			return
		}
	}
	fmt.Println("no")
}
