package leetcode

import "testing"

func TestMinOperations1758(t *testing.T) {
	a := "0100"
	if r := minOperation1758(a); r != 1 {
		t.Fatalf("expect 1 get %d", 1)
	}
	if r := minOperations1758_2(a); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	a = "0"
	if r := minOperation1758(a); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	if r := minOperations1758_2(a); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	a = "10"
	if r := minOperation1758(a); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	if r := minOperations1758_2(a); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	a = "1111"
	if r := minOperation1758(a); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := minOperations1758_2(a); r != 2 {
		t.Fatalf("expect 0 get %d", r)
	}
}
