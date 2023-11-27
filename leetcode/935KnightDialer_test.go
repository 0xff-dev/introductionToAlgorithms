package leetcode

import "testing"

func TestKnightDialer(t *testing.T) {
	if r := knightDialer(1); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	if r := knightDialer(2); r != 20 {
		t.Fatalf("expect 20 get %d", r)
	}
	if r := knightDialer(3); r != 46 {
		t.Fatalf("expect 46 get %d", r)
	}
	if r := knightDialer(20); r != 58689536 {
		t.Fatalf("expect 58689536 get %d", r)
	}
	if r := knightDialer(3131); r != 136006598 {
		t.Fatalf("expect 136006598 get %d", r)
	}
}
