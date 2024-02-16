package leetcode

import "testing"

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	arr := []int{5, 5, 1}
	k := 1
	exp := 1
	if r := findLeastNumOfUniqueInts(arr, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	arr = []int{4, 3, 1, 1, 3, 3, 2}
	k = 3
	exp = 2
	if r := findLeastNumOfUniqueInts(arr, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
