package leetcode

import "testing"

func TestContractor2336(t *testing.T) {
	c := Constructor2336()
	c.AddBack(2)
	if r := c.PopSmallest(); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := c.PopSmallest(); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	if r := c.PopSmallest(); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	c.AddBack(1)
	if r := c.PopSmallest(); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	if r := c.PopSmallest(); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	if r := c.PopSmallest(); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
