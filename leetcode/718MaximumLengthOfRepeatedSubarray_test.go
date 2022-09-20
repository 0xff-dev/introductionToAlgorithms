package leetcode

import "testing"

func TestFindLength(t *testing.T) {
	n1, n2 := []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}
	if r := findLength(n1, n2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	n1, n2 = []int{0, 0, 0}, []int{0, 0, 0}
	if r := findLength(n1, n2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	n1, n2 = []int{1, 0, 0, 0, 1, 0, 0, 1, 0, 0}, []int{0, 1, 1, 1, 0, 1, 1, 1, 0, 0}
	if r := findLength(n1, n2); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
