package leetcode

import (
	"reflect"
	"testing"
)

func TestFindAllPeople(t *testing.T) {
	n := 6
	meetings := [][]int{
		{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
	}
	firstPerson := 1
	exp := []int{0, 1, 2, 3, 5}

	if r := findAllPeople(n, meetings, firstPerson); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 4
	meetings = [][]int{
		{3, 1, 3}, {1, 2, 2}, {0, 3, 3},
	}
	firstPerson = 3
	exp = []int{0, 1, 3}
	if r := findAllPeople(n, meetings, firstPerson); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n = 5
	meetings = [][]int{
		{3, 4, 2}, {1, 2, 1}, {2, 3, 1},
	}
	firstPerson = 1
	exp = []int{0, 1, 2, 3, 4}
	if r := findAllPeople(n, meetings, firstPerson); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
