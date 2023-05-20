package leetcode

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string{
		{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"},
	}
	exp := []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}
	if r := calcEquation(equations, values, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
