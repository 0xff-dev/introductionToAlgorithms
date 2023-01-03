package leetcode

import "testing"

func TestMinDeletionSize(t *testing.T) {
	strs := []string{"cba", "daf", "ghi"}
	if r := minDeletionSize(strs); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	strs = []string{"a", "b"}
	if r := minDeletionSize(strs); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	strs = []string{"zyx", "wvu", "tsr"}
	if r := minDeletionSize(strs); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
