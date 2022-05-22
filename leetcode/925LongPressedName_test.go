package leetcode

import "testing"

func TestIsLongPressedName(t *testing.T) {
	name, typed := "alex", "aaleex"
	if !isLongPressedName(name, typed) {
		t.Fatalf("expect true get false")
	}
	if !isLongPressName2(name, typed) {
		t.Fatalf("expect true get false")
	}

	name, typed = "alex", "alex"
	if !isLongPressedName(name, typed) {
		t.Fatalf("expect true get fales")
	}
	if !isLongPressName2(name, typed) {
		t.Fatalf("expect true get false")
	}

	name, typed = "saeed", "ssaaedd"
	if isLongPressedName(name, typed) {
		t.Fatalf("expect false get true")
	}
	if isLongPressName2(name, typed) {
		t.Fatalf("expect false get true")
	}

	name, typed = "laiden", "laiden"
	if !isLongPressedName(name, typed) {
		t.Fatalf("expect true get false")
	}
	if !isLongPressName2(name, typed) {
		t.Fatalf("expect tru get false")
	}

	name, typed = "alex", "aaleexa"
	if isLongPressedName(name, typed) {
		t.Fatalf("expect false get true")
	}

	if isLongPressName2(name, typed) {
		t.Fatalf("expect false get true")
	}
}
