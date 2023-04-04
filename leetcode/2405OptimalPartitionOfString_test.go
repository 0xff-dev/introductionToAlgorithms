package leetcode

import "testing"

func TestPartitionString(t *testing.T) {
	s := "abacaba"
	if r := partitionString(s); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	s = "ssssss"
	if r := partitionString(s); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
