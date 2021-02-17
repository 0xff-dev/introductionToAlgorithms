package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	if r := maxArea(input); r != 49 {
		t.Fatalf("expect 49 get %d", r)
	}
	input = []int{1, 1}
	if r := maxArea(input); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	input = []int{4, 3, 2, 1, 4}
	if r := maxArea(input); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}

	input = []int{1, 2, 1}
	if r := maxArea(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
