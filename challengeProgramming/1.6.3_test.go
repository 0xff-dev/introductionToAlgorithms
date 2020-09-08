package challengeProgramming

import "testing"

func TestFindSum(t *testing.T) {
	aim := 4
	nums := []int{1, 2, 3, 4}
	if !FindSum(aim, nums) {
		t.Fatalf("can find")
	}
}
