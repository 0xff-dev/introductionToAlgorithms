package leetcode

import "testing"

func TestRestoreString(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	if r := restoreString(s, indices); r != "leetcode" {
		t.Fatalf("expect 'leetcode' get '%s'", r)
	}
}
