package successful_pairs_of_spells_and_potions

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n := len(spells)
	m := len(potions)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		// 二分查找
		l, r, mid := 0, m-1, 0
		for l <= r {
			mid = (l + r) >> 1
			// 如果大于success，继续探索边界，否则缩小边界
			if int64(spells[i])*int64(potions[mid]) >= success {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		ans[i] = m - l
	}
	return ans
}
