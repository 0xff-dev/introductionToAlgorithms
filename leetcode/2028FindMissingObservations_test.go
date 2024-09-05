package leetcode

import (
	"reflect"
	"testing"
)

func TestMissingRolls(t *testing.T) {
	rolls, mean, n, exp := []int{3, 2, 4, 3}, 4, 2, []int{6, 6}
	if r := missingRolls(rolls, mean, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	rolls, mean, n, exp = []int{1, 5, 6}, 3, 4, []int{3, 2, 2, 2}
	if r := missingRolls(rolls, mean, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	rolls, mean, n, exp = []int{1, 2, 3, 4}, 6, 4, []int{}
	if r := missingRolls(rolls, mean, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
