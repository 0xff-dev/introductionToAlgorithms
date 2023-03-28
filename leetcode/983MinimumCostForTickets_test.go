package leetcode

import "testing"

func TestMincostTickets(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	if r := mincostTickets(days, costs); r != 11 {
		t.Fatalf("expect 11 get %d", r)
	}
}
