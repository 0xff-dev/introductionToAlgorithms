package leetcode

import "testing"

func TestReadBinaryWatch(t *testing.T) {
	n := 1
	r := readBinaryWatch(n)
	t.Logf("%v", r)

	t.Logf("%v", readBinaryWatch(9))
}
