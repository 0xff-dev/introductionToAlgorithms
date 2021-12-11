package leetcode

func singleNumber1(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}
	return a
}
