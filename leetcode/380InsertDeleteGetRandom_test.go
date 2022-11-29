package leetcode

import "testing"

func TestConstructor380(t *testing.T) {
	c := Constructor380()
	if !c.Insert(1) {
		t.Fatalf("insert 1 expect true get false")
	}

	if c.Remove(2) {
		t.Fatalf("remove 2 expect false get true")
	}

	if !c.Insert(2) {
		t.Fatalf("insert 2 expect true get false")
	}

	r := c.GetRandom()
	if r != 1 && r != 2 {
		t.Fatalf("GetRandom expect 1 or 2, get: %d", r)
	}

	if !c.Remove(1) {
		t.Fatalf("remove 1 expect true get false")
	}

	if c.Insert(2) {
		t.Fatalf("insert 2 expect false get true")
	}

	r = c.GetRandom()
	if r != 1 && r != 2 {
		t.Fatalf("GetRandom expect 1 or 2 get %d", r)
	}
}
