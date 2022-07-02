package leetcode

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	stack := make([]int, 0)
	for i := 0; i <= len(nums); i++ {
		combination(0, i, nums, &stack, &res)
	}
	return res
}

func combination(startIndex, length int, nums []int, stack *[]int, res *[][]int) {
	if length == 0 {
		tmp := make([]int, len(*stack))
		copy(tmp, *stack)
		*res = append(*res, tmp)
		return
	}
	if startIndex >= len(nums) {
		return
	}
	*stack = append(*stack, nums[startIndex])
	combination(startIndex+1, length-1, nums, stack, res)
	*stack = (*stack)[:len(*stack)-1]
	combination(startIndex+1, length, nums, stack, res)
}
