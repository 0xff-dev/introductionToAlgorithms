package leetcode

import "testing"

func TestMinOperations2870(t *testing.T) {
	n := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	if r := minOperations2870(n); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	n = []int{2, 1, 2, 2, 3, 3}
	if r := minOperations2870(n); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	n = []int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}
	if r := minOperations2870(n); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
