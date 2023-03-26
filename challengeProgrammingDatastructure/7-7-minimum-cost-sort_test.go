package challengeprogrammingdatastructure

import "testing"

func TestMinimumCostSort(t *testing.T) {
	nums := []int{1, 5, 3, 4, 2}
	if r := MinimumCostSort(nums); r != 7 {
		t.Fatalf("expect 7 get %d", r)
	}
}
