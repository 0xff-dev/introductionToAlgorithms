package leetcode

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	n := []int{1, 2, 3, 1, 1, 3}
	if r := numIdenticalPairs(n); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	n = []int{1, 1, 1, 1}
	if r := numIdenticalPairs(n); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	n = []int{1, 2, 3}
	if r := numIdenticalPairs(n); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
