package leetcode

import "testing"

func TestConbine(t *testing.T) {
	n, k := 4, 2
	r := combine(n, k)
	t.Logf("%v", r)

	n, k = 1, 1
	t.Logf("%v", combine(n, k))

}
