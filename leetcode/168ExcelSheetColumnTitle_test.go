package leetcode

import "testing"

func TestConvertToTitle(t *testing.T) {
	if r := convertToTitle(1); r != "A" {
		t.Fatalf("expect 'A' get %s", r)
	}
	if r := convertToTitle(28); r != "AB" {
		t.Fatalf("expect 'AB' get %s", r)
	}
	if r := convertToTitle(701); r != "ZY" {
		t.Fatalf("expect 'ZY' get %s", r)
	}
}
