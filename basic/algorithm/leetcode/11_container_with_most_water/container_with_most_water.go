package container_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		res = max(min(height[left], height[right])*(right-left), res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
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
