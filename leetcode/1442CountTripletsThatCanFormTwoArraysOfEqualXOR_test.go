package leetcode

import "testing"

func TestCountTriplets(t *testing.T) {
	arr := []int{2, 3, 1, 6, 7}
	exp := 4
	if r := countTriplets(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	arr = []int{1, 1, 1, 1, 1}
	exp = 10
	if r := countTriplets(arr); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
