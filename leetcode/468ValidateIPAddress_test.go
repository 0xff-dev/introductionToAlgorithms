package leetcode

import "testing"

func TestValidIPAdress(t *testing.T) {
	a := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	t.Logf(validIPAddress(a))
	t.Logf(validIPAddress("::1"))
}
