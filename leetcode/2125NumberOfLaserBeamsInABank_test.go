package leetcode

import "testing"

func TestNumberOfBeams(t *testing.T) {
	bank := []string{
		"011001", "000000", "010100", "001000",
	}
	if r := numberOfBeams(bank); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	bank = []string{
		"000", "111", "000",
	}
	if r := numberOfBeams(bank); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
