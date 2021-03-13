package leetcode

import "testing"

func TestRestoreIpAddresses(t *testing.T) {
	input := "0000"
	r := restoreIpAddresses(input)
	t.Logf("%v", r)
	input = "25525511135"
	r = restoreIpAddresses(input)
	t.Logf("%v", r)

	input = "1111"
	r = restoreIpAddresses(input)
	t.Logf("%v", r)

	input = "010010"
	r = restoreIpAddresses(input)
	t.Logf("%v", r)

	input = "101023"
	r = restoreIpAddresses(input)
	t.Logf("%v", r)

	input = "0690"
	r = restoreIpAddresses(input)
	t.Logf("%v", r)

	input = "1980"
	r = restoreIpAddresses(input)
	t.Logf("%v", r)
}
