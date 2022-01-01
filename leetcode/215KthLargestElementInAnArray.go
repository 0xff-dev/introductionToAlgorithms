package leetcode

func findKthLargest(nums []int, k int) int {
	start, end := 0, len(nums)-1
	for {
		i := quickSelect(nums, start, end)
		if i == k-1 {
			return nums[i]
		}
		if i < k {
			start = i + 1
			continue
		}
		end = i - 1
	}
}

func quickSelect(nums []int, start, end int) int {
	if start == end {
		return start
	}
	cmp := nums[start]
	cmpIdx := start
	for walker := start + 1; walker <= end; walker++ {
		if nums[walker] > cmp {
			cmpIdx++
			nums[walker], nums[cmpIdx] = nums[cmpIdx], nums[walker]
		}
	}

	nums[start], nums[cmpIdx] = nums[cmpIdx], nums[start]
	return cmpIdx
}
