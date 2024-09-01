package leetcode

import "sort"

func wiggleSort(nums []int) {
	l := len(nums)
	tmp := make([]int, l)
	copy(tmp, nums)
	sort.Ints(tmp)
	left, right := l/2, l-1
	if l&1 == 0 {
		left--
	}
	end := left
	index := 0
	for right > end {
		nums[index] = tmp[left]
		nums[index+1] = tmp[right]
		right, left = right-1, left-1
		index += 2
	}
	if left >= 0 {
		nums[index] = tmp[left]
	}
}
