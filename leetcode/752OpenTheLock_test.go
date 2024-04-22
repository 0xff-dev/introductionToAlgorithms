package leetcode

import "testing"

func TestOpenLock(t *testing.T) {
	d := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	exp := 6
	if r := openLock(d, target); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	d = []string{"8888"}
	target = "0009"
	exp = 1
	if r := openLock(d, target); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
