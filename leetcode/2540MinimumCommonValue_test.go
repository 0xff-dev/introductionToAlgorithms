package leetcode

import "testing"

func TestGetCommon(t *testing.T) {
	n1, n2 := []int{1, 2, 3}, []int{2, 4}
	if r := getCommon(n1, n2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	n1, n2 = []int{1, 2, 3, 6}, []int{2, 3, 4, 5}
	if r := getCommon(n1, n2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
