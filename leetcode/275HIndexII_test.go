package leetcode

import "testing"

func TestHIndex(t *testing.T) {
	citations := []int{0, 1, 3, 5, 6}
	exp := 3
	if r := hIndex(citations); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	citations = []int{1, 2, 100}
	exp = 2
	if r := hIndex(citations); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
