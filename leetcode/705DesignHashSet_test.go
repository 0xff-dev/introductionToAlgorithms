package leetcode

import "testing"

func TestConstructor705(t *testing.T) {
	c := Constructor705()
	c.Add(1)
	c.Add(2)
	if !c.Contains(1) {
		t.Fatalf("expect true get false")
	}
	if c.Contains(3) {
		t.Fatalf("expect false get true")
	}
	c.Add(2)
	if !c.Contains(2) {
		t.Fatalf("expect true get false")
	}
	c.Remove(2)
	if c.Contains(2) {
		t.Fatalf("expect false get true")
	}
}
