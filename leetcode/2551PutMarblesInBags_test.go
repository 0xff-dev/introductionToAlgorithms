package leetcode

import "testing"

func TestPutMarbles(t *testing.T) {
	w := []int{1, 2, 5, 1}
	k := 2
	if r := putMarbles(w, k); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	w = []int{1, 3}
	k = 2
	if r := putMarbles(w, k); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
