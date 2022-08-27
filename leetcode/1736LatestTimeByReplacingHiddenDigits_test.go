package leetcode

import "testing"

func TestMaximumTime(t *testing.T) {
	time := "2?:?0"
	if r := maximumTime(time); r != "23:50" {
		t.Fatalf("expect %s get %s", "23:50", r)
	}

	time = "0?:3?"
	if r := maximumTime(time); r != "09:39" {
		t.Fatalf("expect %s get %s", "09:39", r)
	}

	time = "1?:22"
	if r := maximumTime(time); r != "19:22" {
		t.Fatalf("expext %s get %s", "19:22", r)
	}

	time = "?0:15"
	if r := maximumTime(time); r != "20:15" {
		t.Fatalf("expect 20:15 get %s", r)
	}
}
