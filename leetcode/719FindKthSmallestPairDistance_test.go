package leetcode

import "testing"

func TestSmallestDistnacePair(t *testing.T) {
	nums, k, exp := []int{1, 3, 1}, 1, 0
	if r := smallestDistancePair(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{1, 1, 1}, 2, 0
	if r := smallestDistancePair(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{1, 6, 1}, 3, 5
	if r := smallestDistancePair(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{62, 100, 4}, 2, 58
	if r := smallestDistancePair(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	nums, k, exp = []int{11, 234, 343, 2, 38, 99, 62, 378, 10, 15, 178, 78, 99, 999, 10101, 40302, 44}, 5, 6
	if r := smallestDistancePair(nums, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
