package leetcode

import "testing"

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	if r := countNumbersWithUniqueDigits(2); r != 91 {
		t.Fatalf("expect 91 get %d", r)
	}

	if r := countNumbersWithUniqueDigits(3); r != 739 {
		t.Fatalf("expect 739 get %d", 4)
	}
}
