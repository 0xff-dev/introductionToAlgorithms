package leetcode

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k, threshold := 3, 4
	if r := numOfSubarrays(arr, k, threshold); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	arr = []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}
	k, threshold = 3, 5
	if r := numOfSubarrays(arr, k, threshold); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
