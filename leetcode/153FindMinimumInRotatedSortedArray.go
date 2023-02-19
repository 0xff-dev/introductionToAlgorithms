package leetcode

func findMin(nums []int) int {
	// 0,1,2,3,4
	// 3,4,0,1,2
	// 4,5,6,7,0,1,2  0-- 6
	// 2, 1
	//4,5,6,7,0
	// 0, 1,2,3,4
	l, r := 0, len(nums)-1
	ans := nums[0]
	for l < r {
		mid := (r-l)/2 + l
		pre := mid - 1
		if pre >= 0 && nums[mid] < nums[pre] {
			ans = nums[mid]
			break
		}
		next := mid + 1
		if next < len(nums) && nums[mid] > nums[next] {
			ans = nums[next]
			break
		}
		if nums[mid] > nums[r] {
			l = mid + 1
			continue
		}
		r = mid - 1
	}
	return ans
}
