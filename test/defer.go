package main

import "fmt"

const (
	x = iota
	_
	y
	z = "pi"
	k
	p = iota
	q
)

type user struct {
	Name string
}

func main() {
	a := &user{
		Name: "张三",
	}
	var b interface{} = *a
	tmp := b.(user)
	tmp.Name = "李四"
	fmt.Println(a.Name)
}
