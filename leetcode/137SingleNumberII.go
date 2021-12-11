package leetcode

func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = (one ^ num) & (^two)
		two = (^one) & (two ^ num)
	}
	return one
}
