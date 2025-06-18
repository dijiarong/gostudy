package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 转换成求最小的值
	l1, l2 := len(nums1), len(nums2)
	k := (l1 + l2) / 2
	if (l1+l2)%2 != 0 {
		return float64(findKSmall(nums1, nums2, k))
	} else {
		return float64(findKSmall(nums1, nums2, k)+findKSmall(nums1, nums2, k+1)) / 2
	}
}

func findKSmall(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		curIndex1 := min(len(nums1), index1+half) - 1
		curIndex2 := min(len(nums2), index2+half) - 1
		val1, val2 := nums1[curIndex1], nums2[curIndex2]
		if val1 <= val2 {
			k -= (curIndex1 - index1 + 1)
			index1 = curIndex1 + 1
		} else {
			k -= (curIndex2 - index2 + 1)
			index2 = curIndex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
