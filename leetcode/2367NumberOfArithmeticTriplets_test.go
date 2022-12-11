package leetcode

import "testing"

func TestArithmeticTriplets(t *testing.T) {
	nums := []int{0, 1, 4, 6, 7, 10}
	diff := 3
	if r := arithmeticTriplets(nums, diff); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	nums = []int{4, 5, 6, 7, 8, 9}
	diff = 2
	if r := arithmeticTriplets(nums, diff); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
