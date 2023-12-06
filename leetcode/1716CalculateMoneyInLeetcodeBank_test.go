package leetcode

import "testing"

func TestTotalMoney(t *testing.T) {
	if r := totalMoney(4); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	if r := totalMoney(10); r != 37 {
		t.Fatalf("expect 37 get %d", r)
	}
	if r := totalMoney(20); r != 96 {
		t.Fatalf("expect 96 get %d", r)
	}
}
