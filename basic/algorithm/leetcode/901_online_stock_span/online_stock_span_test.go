package online_stock_span

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {
	stockSpanner := Constructor()
	fmt.Println(stockSpanner.Next(100))
	fmt.Println(stockSpanner.Next(80))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(70))
	fmt.Println(stockSpanner.Next(60))
	fmt.Println(stockSpanner.Next(75))
	fmt.Println(stockSpanner.Next(85))
}
