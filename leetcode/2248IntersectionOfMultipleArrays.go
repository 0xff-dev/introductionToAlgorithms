package leetcode

func intersection2248(nums [][]int) []int {
	length := len(nums)
	ans := make([]int, 0)
	set := make([]int, 1001)
	for idx := 0; idx < length; idx++ {
		for _, n := range nums[idx] {
			set[n]++
		}
	}
	for n := 1; n < 1001; n++ {
		if set[n] == length {
			ans = append(ans, n)
		}
	}

	return ans
}
