package leetcode

import "testing"

func TestMinPartitions(t *testing.T) {
	n := "32"
	if r := minPartitions(n); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	n = "27346209830709182346"
	if r := minPartitions(n); r != 9 {
		t.Fatalf("expect 9 get %d", r)
	}
}
