package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			i, j := mid, mid
			for ; i >= 0 && nums[i] == target; i-- {
			}
			for ; j < len(nums) && nums[j] == target; j++ {
			}
			return []int{i + 1, j - 1}
		}
		if nums[mid] < target {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return []int{-1, -1}
}
