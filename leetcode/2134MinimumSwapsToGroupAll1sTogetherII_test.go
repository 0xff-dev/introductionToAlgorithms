package leetcode

import (
	"testing"
)

func TestMinSwaps(t *testing.T) {
	s, exp := []int{0, 1, 0, 1, 1, 0, 0}, 1
	if r := minSwaps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = []int{0, 1, 1, 1, 0, 0, 1, 1, 0}, 2
	if r := minSwaps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	s, exp = []int{1, 1, 0, 0, 1}, 0
	if r := minSwaps(s); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
