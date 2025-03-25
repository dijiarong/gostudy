package my_calendar1

import (
	"fmt"
	"testing"
)

func TestMy(t *testing.T) {
	c := Constructor()
	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(15, 25))
	fmt.Println(c.Book(20, 30))
}
