package leetcode

func sortedSquares(nums []int) []int {
	index := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			index = i
			break
		}
	}
	i, j := index-1, index
	result := make([]int, len(nums))
	index = 0
	for ; i >= 0 && j < len(nums); index++ {
		a := nums[i] * nums[i]
		b := nums[j] * nums[j]
		if a < b {
			result[index] = a
			i--
			continue
		}
		result[index] = b
		j++
	}
	for ; i >= 0; i, index = i-1, index+1 {
		result[index] = nums[i] * nums[i]
	}
	for ; j < len(nums); j, index = j+1, index+1 {
		result[index] = nums[j] * nums[j]
	}
	return result
}
