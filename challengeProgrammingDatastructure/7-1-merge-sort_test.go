package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{8, 5, 9, 2, 6, 3, 7, 1, 10, 4}
	r := MergeSort(nums, 0, 9)
	exp1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exp2 := 34
	if r != exp2 || !reflect.DeepEqual(nums, exp1) {
		t.Fatalf("expect diff %d get %d, expect nums %v get %v", exp2, r, exp1, nums)
	}
}
