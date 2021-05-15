package leetcode

import "testing"

func TestMaxUncrossedLines(t *testing.T) {
	a, b := []int{1, 4, 2}, []int{1, 2, 4}
	if r := maxUncrossedLines(a, b); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	a, b = []int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}
	if r := maxUncrossedLines(a, b); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	a, b = []int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}
	if r := maxUncrossedLines(a, b); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
