package leetcode

import "testing"

func TestConstructor677(t *testing.T) {
	c := Constructor677()
	c.Insert("apple", 3)
	if r := c.Sum("ap"); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
	c.Insert("app", 2)
	if r := c.Sum("ap"); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
