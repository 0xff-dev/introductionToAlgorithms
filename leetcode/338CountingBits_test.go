package leetcode

import "testing"

func TestCountBits(t *testing.T) {
	r := countBits(2)
	t.Logf("%v ", r)
	r = countBits(5)
	t.Logf("%v ", r)
}
