package leetcode

import "testing"

func TestFlipLights(t *testing.T) {
	n, k := 1, 1
	exp := 2
	if r := flipLights(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	n, k = 2, 1
	exp = 3
	if r := flipLights(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	n, k = 3, 1
	exp = 4
	if r := flipLights(n, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
