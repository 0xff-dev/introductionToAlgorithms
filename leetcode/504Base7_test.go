package leetcode

import "testing"

func TestConvertToBase7(t *testing.T) {
	if r := convertToBase7(100); r != "202" {
		t.Fatalf("expect 202 get %s", r)
	}

	if r := convertToBase7(-7); r != "-10" {
		t.Fatalf("expect -10 get %s", r)
	}

	if r := convertToBase7(0); r != "0" {
		t.Fatalf("expect 0 get %s", r)
	}
}
