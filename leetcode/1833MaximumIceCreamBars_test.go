package leetcode

import "testing"

func TestMaxIceCream(t *testing.T) {
	costs, coins := []int{1, 3, 2, 4, 1}, 7
	if r := maxIceCream(costs, coins); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	costs, coins = []int{10, 6, 8, 7, 7, 8}, 5
	if r := maxIceCream(costs, coins); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	costs, coins = []int{1, 6, 3, 1, 2, 5}, 20
	if r := maxIceCream(costs, coins); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
