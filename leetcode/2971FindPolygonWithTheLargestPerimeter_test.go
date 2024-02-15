package leetcode

import "testing"

func TestLargestPerimeter(t *testing.T) {
	nums := []int{5, 5, 5}
	if r := largestPerimeter(nums); r != 15 {
		t.Fatalf("expect 15 get %d", r)
	}
	nums = []int{1, 12, 1, 2, 5, 50, 3}
	if r := largestPerimeter(nums); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
	nums = []int{5, 5, 50}
	if r := largestPerimeter(nums); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
