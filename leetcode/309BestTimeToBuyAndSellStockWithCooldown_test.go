package leetcode

import "testing"

func TestMaxProfit309(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	if r := maxProfit309(prices); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	prices = []int{1}
	if r := maxProfit309(prices); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
