package offer

import "testing"

func TestMinNumberInRotateArray(t *testing.T) {
	input := []int{3,4,5,1,2}
	if MinNumberInRotateArray(input) != 1 {
		t.Fatalf("except 1")
	}
}
