package range_sum_query_immutable

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	tmp := make([]int, len(nums)+1)
	for i, v := range nums {
		tmp[i+1] = tmp[i] + v
	}
	return NumArray{nums: tmp}
}

func (self *NumArray) SumRange(left int, right int) int {
	return self.nums[right+1] - self.nums[left]
}
