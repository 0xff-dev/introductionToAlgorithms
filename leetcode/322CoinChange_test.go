package leetcode

import "testing"

func TestCoinChange(t *testing.T) {
	coins, amount := []int{1, 2, 5}, 11
	if r := coinChange(coins, amount); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	coins, amount = []int{2}, 3
	if r := coinChange(coins, amount); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	coins, amount = []int{1}, 0
	if r := coinChange(coins, amount); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
