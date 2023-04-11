package leetcode

import "testing"

func TestRemoveStars(t *testing.T) {
	s := "leet**cod*e"
	if r := removeStars(s); r != "lecoe" {
		t.Fatalf("expect 'lecoe' get %s", r)
	}
}
