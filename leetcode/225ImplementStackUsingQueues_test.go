package leetcode

import "testing"

func TestConstroctor225(t *testing.T) {
	s := Constructor225()
	s.Push(1)
	s.Push(2)
	if r := s.Top(); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := s.Pop(); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if s.Empty() {
		t.Fatalf("expect false get true")
	}
}
