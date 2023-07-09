package leetcode

import "testing"

func TestLargestVariance(t *testing.T) {
	s := "aababbb"
	if r := largestVariance(s); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
