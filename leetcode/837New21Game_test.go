package leetcode

import "testing"

func TestNew21Game(t *testing.T) {
	n, k, maxPts := 10, 1, 10
	t.Logf("ans: %.5f\n", new21Game(n, k, maxPts))
}
