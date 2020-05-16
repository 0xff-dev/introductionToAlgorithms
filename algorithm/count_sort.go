package algorith

func maxNum(nums []int) int {
	maxInt := nums[0]
	for index := 1; index < len(nums); index++ {
		if nums[index] > maxInt {
			maxInt = nums[index]
		}
	}
	return maxInt
}

func CountSort(nums []int) []int {
	length := len(nums)
	maxInt := maxNum(nums)
	countArray := make([]int, maxInt+1)
	for index := 1; index < length; index++ {
		countArray[nums[index]]++
	}
	for index := 1; index <= maxInt; index++ {
		countArray[index] = countArray[index-1] + countArray[index]
	}
	// 倒序是为了保证是稳定的排序，如果正向遍历，数字刚好全反过来
	result := make([]int, length)
	for index := length - 1; index >= 0; index-- {
		result[countArray[nums[index]]] = nums[index]
		countArray[nums[index]]--
	}
	return result
}
