package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	target, n := []int{1, 3}, 3
	exp := []string{"Push", "Push", "Pop", "Push"}
	if r := buildArray(target, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	target, n = []int{1, 2, 3}, 3
	exp = []string{"Push", "Push", "Push"}
	if r := buildArray(target, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	target, n = []int{1, 2}, 4
	exp = []string{"Push", "Push"}
	if r := buildArray(target, n); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
