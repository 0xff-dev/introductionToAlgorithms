package leetcode

import "testing"

func TestGenerat(t *testing.T) {
	for n := 1; n < 10; n++ {
		t.Logf("n=%d --> %v", n, generate(n))
	}
}
