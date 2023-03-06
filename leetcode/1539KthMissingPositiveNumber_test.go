package leetcode

import "testing"

func TestFindKthPositive(t *testing.T) {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	if r := findKthPositive(arr, k); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
	arr = []int{1, 2, 3, 4}
	k = 2
	if r := findKthPositive(arr, k); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
