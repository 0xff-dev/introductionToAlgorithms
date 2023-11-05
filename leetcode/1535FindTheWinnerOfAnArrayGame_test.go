package leetcode

import "testing"

func TestGetWinner(t *testing.T) {
	if r := getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
	if r := getWinner([]int{3, 2, 1}, 10); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
