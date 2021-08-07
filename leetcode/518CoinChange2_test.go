package leetcode

import "testing"

func TestChange(t *testing.T) {
	amount, coins := 5, []int{1, 2, 5}
	r := change(amount, coins)
	t.Logf("%v", r)

	amount, coins = 10, []int{10}
	r = change(amount, coins)
	t.Logf("%v", r)
}
