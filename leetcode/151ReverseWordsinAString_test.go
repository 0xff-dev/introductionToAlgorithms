package leetcode

import "testing"

func TestReverseWords(t *testing.T) {
	if r := reverseWords("the sky is blue"); r != "blue is sky the" {
		t.Fatalf("expect 'blue is sky the' get %q", r)
	}
	if r := reverseWords("  hello world  "); r != "world hello" {
		t.Fatalf("expect 'world hello', get %q", r)
	}
	if r := reverseWords("a good   example"); r != "example good a" {
		t.Fatalf("expect example good a get %q", r)
	}
	if r := reverseWords("a good   example"); r != "example good a" {
		t.Fatalf("expect 'example good a' get %q", r)
	}
	if r := reverseWords("Alice does not even like bob"); r != "bob like even not does Alice" {
		t.Fatalf("expect 'bob like even not does Alice' get %q", r)
	}
}
