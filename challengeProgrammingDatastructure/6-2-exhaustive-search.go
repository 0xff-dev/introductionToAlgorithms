package challengeprogrammingdatastructure

// O(2^n)
func exhaustiveSearch(nums []int, index, sum, target, l int) bool {
	if index >= len(nums) || l <= 0 {
		return sum == target
	}
	return exhaustiveSearch(nums, index+1, sum, target, l) || exhaustiveSearch(nums, index+1, sum+nums[index], target, l-1)
}
func ExhaustiveSearch(nums []int, target int) bool {
	for l := 1; l <= len(nums); l++ {
		if exhaustiveSearch(nums, 0, 0, target, l) {
			return true
		}
	}
	return false
}
