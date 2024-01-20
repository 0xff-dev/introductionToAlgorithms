package leetcode

import "testing"

func TestNextGreaterElement556(t *testing.T) {
	n := 12
	if r := nextGreaterElement556(n); r != 21 {
		t.Fatalf("expect 21 get %d", r)
	}
	n = 21
	if r := nextGreaterElement556(n); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	n = 2147483486
	if r := nextGreaterElement556(n); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	n = 34125314
	if r := nextGreaterElement556(n); r != 34125341 {
		t.Fatalf("expect 34125341 get %d", r)
	}
}
