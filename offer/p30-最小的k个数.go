package offer

func GetLeastNumbers(nums []int, k int) []int {
	res := make([]int, k)
	if k == 0 {
		return res
	}
	if k >= len(nums) {
		return nums
	}
	// 1,2,3,4,5最小的三个数
	start, end := 0, len(nums)-1
	index := partition(nums, start, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(nums, start, end)
	}
	for i := 0; i < k; i++ {
		res[i] = nums[i]
	}
	return res
}

func partition(nums []int, start, end int) int {
	data := nums[start]
	location, walker := start, start
	for walker <= end {
		if nums[walker] < data {
			location++
			nums[location], nums[walker] = nums[walker], nums[location]
		}
		walker++
	}
	nums[start], nums[location] = nums[location], nums[start]
	return location
}
