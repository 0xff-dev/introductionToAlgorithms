package leetcode

import "testing"

func TestConstructor432(t *testing.T) {
	c := Constructor432()
	c.Inc("hello")
	c.Inc("hello")
	if r := c.GetMaxKey(); r != "hello" {
		t.Fatalf("expect hello get %s", r)
	}
	if r := c.GetMinKey(); r != "hello" {
		t.Fatalf("expect hello get %s", r)
	}
	c.Inc("leet")
	if r := c.GetMaxKey(); r != "hello" {
		t.Fatalf("expect hello get %s", r)
	}
	if r := c.GetMinKey(); r != "leet" {
		t.Fatalf("expect leet get %s", r)
	}

	c1 := Constructor432()
	c1.Inc("a")
	c1.Inc("b")
	c1.Inc("b")
	c1.Inc("b")
	c1.Inc("b")

	c1.Dec("b")
	c1.Dec("b")

	if r := c1.GetMaxKey(); r != "b" {
		t.Fatalf("expect b get %s", r)
	}
	if r := c1.GetMinKey(); r != "a" {
		t.Fatalf("expect a get %s", r)
	}
}
