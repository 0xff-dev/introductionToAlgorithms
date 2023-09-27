package leetcode

import "testing"

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	if r := decodeString(s); r != "aaabcbc" {
		t.Fatalf("expect aaabcbc get %s", r)
	}
	s = "3[a2[c]]"
	if r := decodeString(s); r != "accaccacc" {
		t.Fatalf("expect accaccacc get %s", r)
	}
	s = "2[abc]3[cd]ef"
	if r := decodeString(s); r != "abcabccdcdcdef" {
		t.Fatalf("expect abcabccdcdcdef get %s", r)
	}
	s = "abc3[cd]xyz"
	if r := decodeString(s); r != "abccdcdcdxyz" {
		t.Fatalf("expect abccdcdcdxyz get %s", r)
	}
	s = "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	if r := decodeString(s); r != "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef" {
		t.Fatalf("expect zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef get %s", r)
	}
}
