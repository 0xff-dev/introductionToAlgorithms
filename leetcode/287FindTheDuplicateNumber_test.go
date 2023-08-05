package leetcode

import "testing"

func TestFindDuplicate(t *testing.T) {
	n := []int{1, 3, 4, 2, 2}
	if r := findDuplicate(n); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	n = []int{3, 1, 3, 4, 2}
	if r := findDuplicate(n); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	n = []int{2, 2, 2, 2, 2}
	if r := findDuplicate(n); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
