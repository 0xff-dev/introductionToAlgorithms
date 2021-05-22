package leetcode

import "testing"

func TestLongestDiverseString(t *testing.T) {
	a, b, c := 1, 1, 7
	r := longestDiverseString(a, b, c)
	t.Logf("%s", r)
	a, b, c = 2, 2, 1
	t.Logf("%s", longestDiverseString(a, b, c))
	a, b, c = 7, 1, 0
	t.Logf("%s", longestDiverseString(a, b, c))
	t.Logf("%s", longestDiverseString(15, 78, 9))
}
