package offer

import "testing"

func TestExistsKinArray(t *testing.T) {
	input := []int{1,2,3,3,3,3,4,5}
	if count := existsKinArray(input, 3); count != 4 {
		t.Fatalf("expect %d get %d",4, count)
	}
}
