package leetcode

import "testing"

func TestComplexNumberMultiply(t *testing.T) {
	n1, n2 := "1+1i", "1+1i"
	exp := "0+2i"
	if r := complexNumberMultiply(n1, n2); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
	n1, n2 = "1+-1i", "1+-1i"
	exp = "0+-2i"
	if r := complexNumberMultiply(n1, n2); r != exp {
		t.Fatalf("expect %s get %s", exp, r)
	}
}
