package leetcode

import "testing"

func TestSubarrayBitWistORs(t *testing.T) {
	a := []int{0}
	if r := subarrayBitwiseORs1(a); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	a = []int{1, 1, 2}
	if r := subarrayBitwiseORs1(a); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	a = []int{1, 2, 4}
	if r := subarrayBitwiseORs1(a); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
