package leetcode

import "testing"

func TestMinOperations1551(t *testing.T) {
	if r := minOperations1551(3); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := minOperations1551(6); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
}
