package leetcode

func search33(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		// 4,5,1,2,3
		if nums[mid] > nums[start] {
			// 递增
			// 4,5,6,1,2
			if target >= nums[start] && target <= nums[mid] {
				end = mid
				continue
			}
			start = mid + 1
			continue
		}
		if target >= nums[mid] && target <= nums[end-1] {
			start = mid + 1
			continue
		}
		end = mid
	}
	return -1
}
