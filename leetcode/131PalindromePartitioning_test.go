package leetcode

import "testing"

func TestPalinePartition(t *testing.T) {
	s := "aab"
	r := partition1(s)
	t.Logf("%v", r)
	s = "aaaa"
	r = partition1(s)
	t.Logf("%v", r)

	t.Logf("%v", partition1(""))
}
