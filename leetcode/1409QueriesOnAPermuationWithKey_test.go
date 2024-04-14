package leetcode

import (
	"reflect"
	"testing"
)

func TestProcessQueries(t *testing.T) {
	queries := []int{3, 1, 2, 1}
	m := 5
	exp := []int{2, 1, 2, 1}
	if r := processQueries(queries, m); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	queries = []int{4, 1, 2, 2}
	m = 4
	exp = []int{3, 1, 2, 0}
	if r := processQueries(queries, m); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	queries = []int{7, 5, 5, 8, 3}
	m = 8
	exp = []int{6, 5, 0, 7, 5}
	if r := processQueries(queries, m); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
