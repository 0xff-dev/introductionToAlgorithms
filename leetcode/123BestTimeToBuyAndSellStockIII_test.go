package leetcode

import "testing"

func TestMaxProfitIII(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	if r := maxProfitIII(prices); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	prices = []int{1, 2, 3, 4, 5}
	if r := maxProfitIII(prices); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	prices = []int{7, 6, 4, 3, 1}
	if r := maxProfitIII(prices); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	prices = []int{2, 1, 2, 0, 1}
	if r := maxProfitIII(prices); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
