package leetcode

import "testing"

func TestMinInsertions(t *testing.T) {
	if r := minInsertions("zzazz"); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	if r := minInsertions("mbadm"); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := minInsertions("leetcode"); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
