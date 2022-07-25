package leetcode

import "testing"

func TestMaxProduct1464(t *testing.T) {
	n := []int{3, 4, 5, 2}
	if r := maxProduct1464(n); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}

	n = []int{3, 7}
	if r := maxProduct1464(n); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}
