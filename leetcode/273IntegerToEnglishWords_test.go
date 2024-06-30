package leetcode

import "testing"

func TestNumberToWords(t *testing.T) {
	num := 123
	exp := "One Hundred Twenty Three"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}

	num = 12345
	exp = "Twelve Thousand Three Hundred Forty Five"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	num = 1234567
	exp = "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	num = 12000
	exp = "Twelve Thousand"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	num = 12345678
	exp = "Twelve Million Three Hundred Forty Five Thousand Six Hundred Seventy Eight"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	num = 1101
	exp = "One Thousand One Hundred One"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	num = 1000000
	exp = "One Million"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}
	num = 111
	exp = "One Hundred Eleven"
	if r := numberToWords(num); r != exp {
		t.Fatalf("expect '%s' get '%s'", exp, r)
	}

}
