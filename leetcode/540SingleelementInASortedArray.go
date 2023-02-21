package leetcode

func singleNonDuplicate(nums []int) int {
	base := 0
	for i, x := range nums {
		base ^= x
		if base != 0 && i&2 == 1 {
			return x
		}
	}
	return -1
}
