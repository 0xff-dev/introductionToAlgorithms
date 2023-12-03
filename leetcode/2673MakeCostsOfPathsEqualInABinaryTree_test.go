package leetcode

import "testing"

func TestMinIncrements(t *testing.T) {
	n := 7
	cost := []int{1, 5, 2, 2, 3, 3, 1}
	if r := minIncrements(n, cost); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	n = 3
	cost = []int{5, 3, 3}
	if r := minIncrements(n, cost); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	n = 15
	cost = []int{1, 5, 2, 2, 3, 3, 1, 1, 2, 3, 4, 5, 6, 7, 8}
	if r := minIncrements(n, cost); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
}
