package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	haystack, needle := "hello", "ll"
	if idx := strStr(haystack, needle); idx != 2 {
		t.Fatalf("expect 2 get %d", idx)
	}

	haystack, needle = "aaaaaaa", "bba"
	if idx := strStr(haystack, needle); idx != -1 {
		t.Fatalf("expect -1 get %d", idx)
	}

	haystack, needle = "", ""
	if idx := strStr(haystack, needle); idx != 0 {
		t.Fatalf("expect 0 get %d", idx)
	}

	haystack, needle = "abc", ""
	if idx := strStr(haystack, needle); idx != 0 {
		t.Fatalf("expect 0 get %d", idx)
	}

	haystack, needle = "abc", " "
	if idx := strStr(haystack, needle); idx != -1 {
		t.Fatalf("expect 0 get %d", idx)
	}

	haystack, needle = " ", " "
	if idx := strStr(haystack, needle); idx != 0 {
		t.Fatalf("expect 0 get %d", idx)
	}

	haystack, needle = "abc", "a"
	if idx := strStr(haystack, needle); idx != 0 {
		t.Fatalf("expect 0 get %d", idx)
	}

	haystack, needle = "          ", "      "
	if idx := strStr(haystack, needle); idx != 0 {
		t.Fatalf("expect 0 get %d", idx)
	}
}
