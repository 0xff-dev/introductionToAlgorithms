package leetcode

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extra := 3
	exp := []bool{true, true, true, false, true}
	if r := kidsWithCandies(candies, extra); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	candies = []int{4, 2, 1, 1, 2}
	extra = 1
	exp = []bool{true, false, false, false, false}
	if r := kidsWithCandies(candies, extra); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
