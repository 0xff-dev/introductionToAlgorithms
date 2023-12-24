package leetcode

import "testing"

func TestValidUtf8(t *testing.T) {
	data := []int{197, 130, 1}
	if !validUtf8(data) {
		t.Fatalf("expect true get false")
	}
	data = []int{235, 140, 4}
	if validUtf8(data) {
		t.Fatalf("expect false get true")
	}
	data = []int{255}
	if validUtf8(data) {
		t.Fatalf("expect false get true")
	}
}
