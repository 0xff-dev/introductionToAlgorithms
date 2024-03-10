package leetcode

import "testing"

func TestPushDominoes(t *testing.T) {
	dominos := "RR.L"
	exp := "RR.L"
	if r := pushDominoes(dominos); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}

	dominos = ".L.R...LR..L.."
	exp = "LL.RR.LLRRLL.."
	if r := pushDominoes(dominos); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
