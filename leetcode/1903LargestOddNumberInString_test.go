package leetcode

import "testing"

func TestLargestOddNumber(t *testing.T) {
	if r := largestOddNumber("52"); r != "5" {
		t.Fatalf("expect 5 get %s", r)
	}
	if r := largestOddNumber("4206"); r != "" {
		t.Fatalf("expect '' get %s", r)
	}
	if r := largestOddNumber("35427"); r != "35427" {
		t.Fatalf("expect 35427 get %s", r)
	}
}
