package leetcode

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	exp := []int{1, 1, 4, 2, 1, 1, 0, 0}

	if r := dailyTemperatures(temperatures); !reflect.DeepEqual(r, exp) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
