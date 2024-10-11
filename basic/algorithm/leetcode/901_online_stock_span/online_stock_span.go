package online_stock_span

import "math"

type pair struct {
	Price int
	Index int
}

type StockSpanner struct {
	Prices   []pair
	CurIndex int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Prices: []pair{{
			Price: math.MaxInt,
			Index: -1,
		}},
		CurIndex: -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	for price >= this.Prices[len(this.Prices)-1].Price {
		this.Prices = this.Prices[:len(this.Prices)-1]
	}
	this.CurIndex++
	this.Prices = append(this.Prices, pair{
		Price: price,
		Index: this.CurIndex,
	})
	return this.CurIndex - this.Prices[len(this.Prices)-2].Index
}
