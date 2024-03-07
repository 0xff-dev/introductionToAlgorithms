package leetcode

import (
	"reflect"
	"testing"
)

func TestAmbiguousCoodinates(t *testing.T) {
	s := "(123)"
	exp := []string{"(1, 23)", "(1, 2.3)", "(12, 3)", "(1.2, 3)"}
	if r := ambiguousCoordinates(s); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
