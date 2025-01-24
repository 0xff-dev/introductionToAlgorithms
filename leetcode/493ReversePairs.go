package leetcode

func merge493(nums []int, start, end int) int {
	mid := start + (end-start)/2
	l, r := start, mid+1
	tmpNums := make([]int, 0)
	count := 0
	for ; l <= mid; l++ {
		for ; r <= end && nums[l] > 2*nums[r]; r++ {

		}
		count += r - 1 - mid
	}
	l, r = start, mid+1
	for l <= mid && r <= end {
		if nums[l] < nums[r] {
			tmpNums = append(tmpNums, nums[l])
			l++
		} else {
			tmpNums = append(tmpNums, nums[r])
			r++
		}
	}
	for ; l <= mid; l++ {
		tmpNums = append(tmpNums, nums[l])
	}
	for ; r <= end; r++ {
		tmpNums = append(tmpNums, nums[r])
	}
	for k := 0; k < len(tmpNums); k++ {
		nums[k+start] = tmpNums[k]
	}
	return count
}

func MergeSort493(nums []int, start, end int) int {
	if start < end {
		mid := start + (end-start)/2
		return MergeSort493(nums, start, mid) + MergeSort493(nums, mid+1, end) + merge493(nums, start, end)
	}
	return 0
}

func reversePairs(nums []int) int {
	return MergeSort493(nums, 0, len(nums)-1)
}
