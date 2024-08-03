package leetcode

import (
	"reflect"
	"testing"
)

func TestAddOperators(t *testing.T) {
	nums, target, exp := "123", 6, []string{"1+2+3", "1*2*3"}

	if r := addOperators(nums, target); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums, target, exp = "232", 8, []string{"2+3*2", "2*3+2"}
	if r := addOperators(nums, target); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	nums, target, exp = "3456237490", 9191, []string{}
	if r := addOperators(nums, target); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
