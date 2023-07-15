package leetcode

import "testing"

func TestDefangIPAddr(t *testing.T) {
	address := "1.1.1.1"
	exp := "1[.]1[.]1[.]1"
	if r := defangIPaddr(address); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
