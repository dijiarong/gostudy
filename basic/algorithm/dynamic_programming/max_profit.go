package dynamicprogramming

import (
	"gostudy/basic/mymath"
)

func MaxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			res = mymath.Math.GetMaxInt(prices[i]-min, res)
		} else {
			min = prices[i]
		}
	}
	return 0
}
