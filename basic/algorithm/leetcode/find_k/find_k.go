package find_k

import "math"

func find_k(arr1, arr2 []int, k int) int {
	l1, l2 := len(arr1), len(arr2)
	k = l1 + l2 - k + 1
	if l1 > l2 {
		return find_k(arr2, arr1, k)
	}
	if l1 == 0 {
		return arr2[k-1]
	}
	if k == 1 {
		return min(arr1[0], arr2[0])
	}
	left, right := max(0, k-l2), min(k, l1)
	for left <= right {
		i := (left + right) / 2 // A中取i个元素
		j := k - i              // B中取j个元素
		start1 := math.MinInt
		if i > 0 {
			start1 = arr1[i-1]
		}
		end1 := math.MaxInt
		if i < l1 {
			end1 = arr1[i]
		}
		start2 := math.MinInt
		if j > 0 {
			start2 = arr2[j-1]
		}
		end2 := math.MaxInt
		if j < l2 {
			end2 = arr2[j]
		}
		if start1 <= end2 && start2 <= end1 {
			return max(start1, start2)
		} else if start1 > end2 {
			right -= 1
		} else {
			left += 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
