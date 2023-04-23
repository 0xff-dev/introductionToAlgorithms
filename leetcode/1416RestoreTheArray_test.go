package leetcode

import "testing"

func TestNumberOfArrays(t *testing.T) {
	s := "1317"
	k := 2000
	if r := numberOfArrays(s, k); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}
	s = "1000"
	k = 10000
	if r := numberOfArrays(s, k); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
	s = "1000"
	k = 10
	if r := numberOfArrays(s, k); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
	s = "2553462832281151811513004352253111"
	k = 455
	if r := numberOfArrays(s, k); r != 21752500 {
		t.Fatalf("expect 21752500 get %d", r)
	}

	s = "1000000000000000000000000000000000"
	k = 10000
	if r := numberOfArrays(s, k); r != 0 {
		t.Fatalf("expect 0 get %d", r)
	}
}
