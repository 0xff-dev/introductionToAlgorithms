package leetcode

import "testing"

func TestMinDistance1(t *testing.T) {
	w1, w2 := "sea", "eat"
	if r := minDistance1(w1, w2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := Solution(w1, w2); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	w1, w2 = "aa", "aa"
	if r := minDistance1(w1, w2); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	if r := Solution(w1, w2); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	w1, w2 = "leetcode", "etco"
	if r := minDistance1(w1, w2); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	if r := Solution(w1, w2); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
