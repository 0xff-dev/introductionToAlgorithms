package offer

import "testing"

func TestCountOfOne(t *testing.T) {
	if CountOfOne(8) != 1 {
		t.Fatalf("except 1")
	}
}
