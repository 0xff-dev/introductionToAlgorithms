package offer

import "testing"

func TestMoreThanHalf(t *testing.T) {
	input := []int{1, 1, 2}
	if MoreThanHalf(input) != 1 {
		t.Fatal("expect 1")
	}
}
