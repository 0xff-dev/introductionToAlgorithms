package leetcode

import "testing"

func TestTribonacci(t *testing.T) {
	if r := tribonacci(4); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	if r := tribonacci(25); r != 1389537 {
		t.Fatalf("expect 1389537 get %d", r)
	}
}
