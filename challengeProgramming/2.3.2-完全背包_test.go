package challengeProgramming

import "testing"

func TestCompleteBakcpack(t *testing.T) {
	weights, values := []int{3, 4, 2}, []int{4, 5, 3}
	if res := CompleteBackpack(weights, values, 7); res != 10 {
		t.Fatalf("expect 10 get %d ", res)
	}
}
