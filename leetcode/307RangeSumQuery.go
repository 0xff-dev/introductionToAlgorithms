package leetcode

// 看解法有线段树..找时间了解一下
type NumArray struct {
	raw []int
}

func Constructor307(nums []int) NumArray {
	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
	}
	return NumArray{raw: append([]int{0}, nums...)}
}

func (this *NumArray) Update(index int, val int) {
	source := this.raw[index+1] - this.raw[index]
	diff := val - source
	for idx := index + 1; idx < len(this.raw); idx++ {
		this.raw[idx] += diff
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.raw[right+1] - this.raw[left]
}
