package leetcode

func permute(nums []int) [][]int {
	r := make([][]int, 0)
	permuteAux(nums, 0, &r)
	return r
}

func permuteAux(nums []int, index int, r *[][]int) {
	if index == len(nums) {
		dest := make([]int, len(nums))
		copy(dest, nums)
		*r = append(*r, dest)
		return
	}
	for idx := index; idx < len(nums); idx++ {
		nums[idx], nums[index] = nums[index], nums[idx]
		permuteAux(nums, index+1, r)
		nums[idx], nums[index] = nums[index], nums[idx]
	}
}
