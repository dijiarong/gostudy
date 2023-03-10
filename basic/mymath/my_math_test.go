package mymath_test

import (
	"fmt"
	"gostudy/basic/mymath"
	"testing"
)

func TestGetNote(t *testing.T) {
	a := 300
	b := 40
	println(mymath.Math.Sub(0, a))
	println(mymath.Math.Sub(0, b))
	println(fmt.Sprintf("%d * %d = %d", a, b, mymath.Math.Multi(a, b)))
	println(fmt.Sprintf("%d / %d = %d", a, b, mymath.Math.Div(a, b)))
	println(string([]rune("<em>")))
	type mmmm struct {
	}
	c := new(mmmm)
	d := new(mmmm)
	println(c)
	println(d)
}
