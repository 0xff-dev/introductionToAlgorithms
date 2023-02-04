package leetcode

import "testing"

func TestCheckInclusion(t *testing.T) {
	s1, s2 := "ab", "eidbaooo"
	if !checkInclusion(s1, s2) {
		t.Fatalf("expect true get false")
	}

	s1, s2 = "ab", "eidboaoo"
	if checkInclusion(s1, s2) {
		t.Fatalf("expect flase get true")
	}
}
