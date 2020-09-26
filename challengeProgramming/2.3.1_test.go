package challengeProgramming

import "testing"

func TestBackpack(t *testing.T) {
	weights, values := []int{2, 1, 3, 2}, []int{2, 3, 4, 2}
	if res := Backpack(weights, values, 5); res != 7 {
		t.Fatalf("expect 7 get %d", res)
	}
}
