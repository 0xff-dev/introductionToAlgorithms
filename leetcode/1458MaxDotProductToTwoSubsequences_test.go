package leetcode

import "testing"

func TestMaxDotProduct(t *testing.T) {
	n1, n2 := []int{2, 1, -2, 5}, []int{3, 0, -6}
	if r := maxDotProduct(n1, n2); r != 18 {
		t.Fatalf("expect 18 get %d", r)
	}
	n1, n2 = []int{3, -2}, []int{2, -6, 7}
	if r := maxDotProduct(n1, n2); r != 21 {
		t.Fatalf("expect 21 get %d", r)
	}
	n1, n2 = []int{-1, -1}, []int{1, 1}
	if r := maxDotProduct(n1, n2); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
}
