package leetcode

import "testing"

func TestPathInZigZagTree(t *testing.T) {
	n := 14
	t.Logf("%v", pathInZigZagTree(n))
	n = 26
	t.Logf("%v", pathInZigZagTree(n))
	n = 1
	t.Logf("%v", pathInZigZagTree(n))
}
