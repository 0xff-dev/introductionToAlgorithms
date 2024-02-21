package leetcode

import "testing"

func TestHIndex274(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	exp := 3
	if r := hIndex274(citations); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	citations = []int{1, 3, 1}
	exp = 1
	if r := hIndex274(citations); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
