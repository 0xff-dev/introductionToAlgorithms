package leetcode

import "testing"

func TestIsAdditiveNumber(t *testing.T) {
	s := "112358"
	if !isAdditiveNumber(s) {
		t.Fatalf("expect true get false")
	}
	s = "199100199"
	if !isAdditiveNumber(s) {
		t.Fatalf("expect true get false")
	}
	s = "123"
	if !isAdditiveNumber(s) {
		t.Fatalf("expect true get false")
	}

	s = "000"
	if !isAdditiveNumber(s) {
		t.Fatalf("expect true get false")
	}

	s = "199001200"
	if isAdditiveNumber(s) {
		t.Fatalf("expect false get true")
	}
	s = "101"
	if !isAdditiveNumber(s) {
		t.Fatalf("expect true get false")
	}
}
func TestAddStr(t *testing.T) {
	s := "1234"
	exp := "64"
	if r := addStr(s, 0, 1, 3); r != exp {
		t.Fatalf("expect' %s' get '%s'", exp, r)
	}
	s = "12345"
	exp = "753"
	if r := addStr(s, 0, 1, 4); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}

	s = "99999"
	//999 + 99
	exp = "8901"
	if r := addStr(s, 0, 2, 4); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
}
