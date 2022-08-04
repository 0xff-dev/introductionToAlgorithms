package leetcode

import "testing"

func TestMirrorReflection(t *testing.T) {
	if r := mirrorReflection(3, 1); r != 1 {
		t.Fatalf("expect 1 get %d", r)
	}
}
