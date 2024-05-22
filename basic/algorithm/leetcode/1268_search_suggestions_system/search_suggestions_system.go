package searchsuggestionssystem

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/search-suggestions-system
/*
给你一个产品数组 products 和一个字符串 searchWord ，products  数组中每个产品都是一个字符串。

请你设计一个推荐系统，在依次输入单词 searchWord 的每一个字母后，
推荐 products 数组中前缀与 searchWord 相同的最多三个产品。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。

请你以二维列表的形式，返回在输入 searchWord 每个字母后相应的推荐产品的列表。
*/
func suggestedProducts(products []string, searchWord string) [][]string {
	// 先排序
	sort.Strings(products)
	res := make([][]string, len(searchWord))
	tmp := ""
	for i, c := range searchWord {
		tmp += string(c)
		curPorducts := []string{}
		for _, product := range products {
			if len(curPorducts) < 3 && strings.HasPrefix(product, tmp) {
				curPorducts = append(curPorducts, product)
			}
			if len(curPorducts) == 3 {
				break
			}
		}
		if len(curPorducts) == 0 && i != 0 {
			curPorducts = res[i-1]
		}
		res[i] = curPorducts
	}
	return res
}
