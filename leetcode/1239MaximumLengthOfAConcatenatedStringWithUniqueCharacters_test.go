package leetcode

import "testing"

func TestMaxLength1239(t *testing.T) {
	arr := []string{"un", "iq", "ue"}
	if r := maxLength1239(arr); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
	arr = []string{"cha", "r", "act", "ers"}
	if r := maxLength1239(arr); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	arr = []string{"abcdefghijklmnopqrstuvwxyz"}
	if r := maxLength1239(arr); r != 26 {
		t.Fatalf("expect 26 get %d", r)
	}
	arr = []string{"abcdefghijklm", "bcdefghijklmn", "cdefghijklmno", "defghijklmnop", "efghijklmnopq", "fghijklmnopqr", "ghijklmnopqrs", "hijklmnopqrst", "ijklmnopqrstu", "jklmnopqrstuv", "klmnopqrstuvw", "lmnopqrstuvwx", "mnopqrstuvwxy", "nopqrstuvwxyz", "opqrstuvwxyza", "pqrstuvwxyzab"}
	if r := maxLength1239(arr); r != 26 {
		t.Fatalf("expect 26 get %d", r)
	}
}
