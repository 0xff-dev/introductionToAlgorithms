package leetcode

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	arr, a, b, c := []int{3, 0, 1, 1, 9, 7}, 7, 2, 3
	if r := countGoodTriplets(arr, a, b, c); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	arr, a, b, c = []int{1, 1, 2, 2, 3}, 0, 0, 1
	if r := countGoodTriplets(arr, a, b, c); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	arr, a, b, c = []int{7, 3, 7, 3, 12, 1, 12, 2, 3}, 5, 8, 1
	if r := countGoodTriplets(arr, a, b, c); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}
