package challengeprogrammingdatastructure

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 1}, {1, 2}, {3, 1}, {3, 4}, {4, 5}, {5, 2},
	}
	exp := []int{0, 3, 1, 4, 5, 2}
	if r := TopologicalSort(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
