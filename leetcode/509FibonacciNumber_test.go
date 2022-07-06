package leetcode

import "testing"

func TestFib(t *testing.T) {
	if r := fib(0); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}

	if r := fib(1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}

	if r := fib(3); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	if r := fib(4); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
