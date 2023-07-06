package leetcode

import "testing"

func TestOrderlyQueue(t *testing.T) {
	s := "gxzv"
	k := 4
	if r := orderlyQueue(s, k); r != "gvxz" {
		t.Fatalf("expect 'gvxz' get %s", r)
	}
}
