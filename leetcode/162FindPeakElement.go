package leetcode

// return it's
func findPeakelement(nums []int) int {
	return findPeakelementAux(nums, 0, len(nums)-1)
}

func findPeakelementAux(nums []int, left, right int) int {
	if left == right {
		return left
	}

	mid := left + (right-left)/2
	if nums[mid] > nums[mid+1] {
		return findPeakelementAux(nums, left, mid)
	}
	return findPeakelementAux(nums, mid+1, right)
}
