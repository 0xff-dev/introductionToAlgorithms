package leetcode

import "testing"

func TestNumWays1269(t *testing.T) {
	steps, arrLen := 3, 2
	if r := numWays1269(steps, arrLen); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	steps, arrLen = 2, 4
	if r := numWays1269(steps, arrLen); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	steps, arrLen = 4, 2
	if r := numWays1269(steps, arrLen); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	steps, arrLen = 430, 148488
	if r := numWays1269(steps, arrLen); r != 525833932 {
		t.Fatalf("expect 525833932 get %d", r)
	}
}
