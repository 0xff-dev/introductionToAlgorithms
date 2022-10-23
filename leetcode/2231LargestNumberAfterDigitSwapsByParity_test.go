package leetcode

import "testing"

func TestLargestInteger(t *testing.T) {
	n := 1234
	if r := largestInteger(n); r != 3412 {
		t.Fatalf("expect %d get %d", 3412, r)
	}

	n = 65875
	if r := largestInteger(n); r != 87655 {
		t.Fatalf("expect 87655 get %d", r)
	}

	n = 247
	if r := largestInteger(n); r != 427 {
		t.Fatalf("expect 427 get %d", r)
	}
}
