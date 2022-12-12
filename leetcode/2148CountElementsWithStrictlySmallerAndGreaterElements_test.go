package leetcode

import "testing"

func TestCountElements(t *testing.T) {
	n := []int{11, 7, 2, 15}
	if r := countElements(n); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
