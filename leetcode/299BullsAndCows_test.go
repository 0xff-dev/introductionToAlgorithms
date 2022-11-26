package leetcode

import "testing"

func TestGetHint(t *testing.T) {
	secret, guess := "1807", "7810"
	if r := getHint(secret, guess); r != "1A3B" {
		t.Fatalf("expect %s get %s", "1A3B", r)
	}

	secret, guess = "1123", "0111"
	if r := getHint(secret, guess); r != "1A1B" {
		t.Fatalf("expect 1A1B get %s", r)
	}

	secret, guess = "1111", "1011"
	if r := getHint(secret, guess); r != "3A0B" {
		t.Fatalf("expect 3A get %s", r)
	}
}
