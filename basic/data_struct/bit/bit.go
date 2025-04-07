package bit

import "fmt"

type Bit struct {
	arr []int
}

func (b *Bit) Put(num int) {
	index := num / 31
	// 首先判断是否超过arr
	for len(b.arr) < index+1 {
		b.arr = append([]int{0}, b.arr...)
	}
	bitIndex := num % 31
	b.arr[index] |= 1
	fmt.Println(bitIndex)
}
