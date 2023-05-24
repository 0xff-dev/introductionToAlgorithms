package leetcode

import "testing"

func TestMaxScore2542(t *testing.T) {
	n1, n2 := []int{1, 3, 3, 2}, []int{2, 1, 3, 4}
	if r := maxScore2542(n1, n2, 3); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}
