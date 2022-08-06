package leetcode

import "testing"

func TestPoorPigs(t *testing.T) {
	b, md, mt := 1000, 15, 60
	if r := poorPigs(b, md, mt); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}

	b, md, mt = 4, 15, 15
	if r := poorPigs(b, md, mt); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	b, md, mt = 4, 15, 30
	if r := poorPigs(b, md, mt); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}
	b, md, mt = 8, 15, 40
	if r := poorPigs(b, md, mt); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	b, md, mt = 5, 15, 60
	if r := poorPigs(b, md, mt); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
