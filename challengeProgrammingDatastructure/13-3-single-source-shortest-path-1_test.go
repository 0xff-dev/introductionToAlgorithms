package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestSingleSourceShortestPath(t *testing.T) {
	table := []adjTable{
		{0, 3, []int{2, 3, 3, 1, 1, 2}},
		{1, 2, []int{0, 2, 3, 4}},
		{2, 3, []int{0, 3, 3, 1, 4, 1}},
		{3, 4, []int{2, 1, 0, 1, 1, 4, 4, 3}},
		{4, 2, []int{2, 1, 3, 3}},
	}
	exp := []int{0, 2, 2, 1, 3}
	ans1, ans2 := SingleSourceShortestPath(5, table)
	if !reflect.DeepEqual(exp, ans1) {
		t.Fatalf("expect %v get %v", exp, ans1)
	}
	if !reflect.DeepEqual(exp, ans2) {
		t.Fatalf("expect %v get %v", exp, ans2)
	}

	/*
			1	3
		0
			2	4
	*/
	table = []adjTable{
		{0, 2, []int{1, 6, 2, 7}},
		{1, 2, []int{3, 5, 2, 8, 4, 4}},
		{2, 1, []int{4, 9}},
		{3, 0, []int{}},
		{4, 0, []int{}},
	}
	ans1, ans2 = SingleSourceShortestPath(5, table)
	exp = []int{0, 6, 7, 11, 10}
	if !reflect.DeepEqual(exp, ans1) {
		t.Fatalf("expect %v get %v", exp, ans1)
	}
	if !reflect.DeepEqual(exp, ans2) {
		t.Fatalf("expect %v get %v", exp, ans2)
	}
}
