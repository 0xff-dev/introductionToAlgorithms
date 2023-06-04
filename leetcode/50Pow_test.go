package leetcode

import "testing"

func TestMyPow(t *testing.T) {
	x, n := 2.0000, 10
	if r := myPow(x, n); r != 1024.0000 {
		t.Fatalf("expect 1024.0000 get %f", r)
	}
}
