package leetcode

import "testing"

func TestIsValidSerialization(t *testing.T) {
	p := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	if !isValidSerialization(p) {
		t.Fatalf("with %s expect true get false", p)
	}
	p = "1,#"
	if isValidSerialization(p) {
		t.Fatalf("with %s expect false get true", p)
	}

	p = "9,#,#,1"
	if isValidSerialization(p) {
		t.Fatalf("with %s expect false get true", p)
	}
	p = "9,#,93,#,9,9,#,#,#"
	if !isValidSerialization(p) {
		t.Fatalf("with %s expect true get false", p)
	}
	p = "9,9,91,#,#,9,#,49,#,#,#"
	if !isValidSerialization(p) {
		t.Fatalf("with %s expect true get false", p)
	}
	p = "8,#,5,#,2,5,#,7,9,#,8,#,#,#,#"
	if !isValidSerialization(p) {
		t.Fatalf("with %s expect true get false", p)
	}
	p = "1,#,#,#,#"
	if isValidSerialization(p) {
		t.Fatalf("with %s expect false get true", p)
	}
	p = "1,#,#"
	if !isValidSerialization(p) {
		t.Fatalf("with %s expect true get false", p)
	}
}
