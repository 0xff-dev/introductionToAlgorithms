package leetcode

import "testing"

func TestMinDeletions1647(t *testing.T) {
	input := "aab"
	if r := minDeletions1647(input); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	input = "aaabbbcc"
	if r := minDeletions1647(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	input = "ceabaacb"
	if r := minDeletions1647(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	input = ""
	if r := minDeletions1647(input); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	input = "aaaaaa"
	if r := minDeletions1647(input); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	input = "aaaaaabbbbbbcccccc"
	if r := minDeletions1647(input); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	input = "abbcccdd"
	if r := minDeletions1647(input); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
}
