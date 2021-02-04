package leetcode

import "testing"

func TestNumTrees(t *testing.T) {
	for i := 1; i <= 19; i++ {
		t.Logf("[%d]=%d", i, numTrees(i))
	}
}
