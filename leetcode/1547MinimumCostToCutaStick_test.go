package leetcode

import "testing"

func TestMinCost(t *testing.T) {
	n := 7
	cuts := []int{1, 3, 4, 5}
	if r := minCost(n, cuts); r != 16 {
		t.Fatalf("expect 16 get %d", r)
	}
	n = 9
	cuts = []int{5, 6, 1, 4, 2}
	if r := minCost(n, cuts); r != 22 {
		t.Fatalf("expect 22 get %d", r)
	}
	n = 2
	cuts = []int{1}
	if r := minCost(n, cuts); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
