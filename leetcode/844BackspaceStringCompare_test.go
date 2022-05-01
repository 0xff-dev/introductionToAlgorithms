package leetcode

import "testing"

func TestBackspaceCompare(t *testing.T) {
	s, t1 := "ab#c", "ad#c"
	if !backspaceCompare(s, t1) {
		t.Fatalf("expect true get false")
	}

	s, t1 = "ab##", "c#d#"
	if !backspaceCompare(s, t1) {
		t.Fatalf("expect ture get false")
	}

	s, t1 = "a#c", "b"
	if backspaceCompare(s, t1) {
		t.Fatalf("expect false get true")
	}

	s, t1 = "###ab", "ab"
	if !backspaceCompare(s, t1) {
		t.Fatalf("expect true get false")
	}

	s, t1 = "a###b", "a#####b"
	if !backspaceCompare(s, t1) {
		t.Fatalf("expect true get false")
	}
}
