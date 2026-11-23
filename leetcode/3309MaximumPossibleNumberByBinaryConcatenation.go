package leetcode

import "math/bits"

func after(x, y int) bool {
	return y<<bits.Len(uint(x))|x > x<<bits.Len(uint(y))|y
}
func maxGoodNumber(nums []int) int {
	a, b, c := nums[0], nums[1], nums[2]
	if after(a, b) {
		a, b = b, a
	}
	if after(a, c) {
		a, c = c, a
	}
	if after(b, c) {
		b, c = c, b
	}

	return a<<bits.Len(uint(b))<<bits.Len(uint(c)) | b<<bits.Len(uint(c)) | c
}
