package leetcode

import "testing"

func TestCountAndSay(t *testing.T) {
	for n := 1; n <= 30; n++ {
		t.Logf("%d -> %s", n, countAndSay(n))
	}
}
