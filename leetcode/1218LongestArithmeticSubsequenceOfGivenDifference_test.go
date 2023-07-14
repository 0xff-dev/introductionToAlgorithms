package leetcode

import "testing"

func TestLongestSubsequence(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	d := 1
	if r := longestSubsequence(arr, d); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	arr = []int{1, 3, 5, 7}
	d = 1
	if r := longestSubsequence(arr, d); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	arr = []int{1, 5, 7, 8, 5, 3, 4, 2, 1}
	d = -2
	if r := longestSubsequence(arr, d); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	arr = []int{7, 7, 7, 7, 7, 7, 7}
	d = 0
	if r := longestSubsequence(arr, d); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
	arr = []int{4, 12, 10, 0, -2, 7, -8, 9, -9, -12, -12, 8, 8}
	d = 0
	if r := longestSubsequence(arr, d); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
