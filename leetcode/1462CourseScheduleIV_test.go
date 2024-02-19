package leetcode

import (
	"reflect"
	"testing"
)

func TestCheckIfPrerequisite(t *testing.T) {
	numCourses := 2
	prerequisties := [][]int{{1, 0}}
	queries := [][]int{{0, 1}, {1, 0}}
	exp := []bool{false, true}
	if r := checkIfPrerequisite(numCourses, prerequisties, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	numCourses = 2
	prerequisties = [][]int{}
	exp = []bool{false, false}
	if r := checkIfPrerequisite(numCourses, prerequisties, queries); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	numCourses = 3
	prerequisties = [][]int{{1, 2}, {1, 0}, {2, 0}}
	queries = [][]int{{1, 0}, {1, 2}}
	exp = []bool{true, true}
}
