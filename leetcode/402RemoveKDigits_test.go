package leetcode

import "testing"

func TestRemoveKdigits(t *testing.T) {
	num := "1432219"
	k := 3
	exp := "1219"
	if r := removeKdigits(num, k); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	num = "10200"
	k = 1
	exp = "200"
	if r := removeKdigits(num, k); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	num = "10"
	k = 2
	exp = "0"
	if r := removeKdigits(num, k); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	num = "101533485478947384734728478473847"
	k = 12
	exp = "13334734728478473847"
	if r := removeKdigits(num, k); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
