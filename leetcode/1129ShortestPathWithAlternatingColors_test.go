package leetcode

import (
	"reflect"
	"testing"
)

func TestShortestAlternatingPaths1129(t *testing.T) {
	n, redEdges, blueEdges := 3, [][]int{{0, 1}, {1, 2}}, [][]int{}
	exp := []int{0, 1, -1}
	if r := shortestAlternatingPaths(n, redEdges, blueEdges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, redEdges, blueEdges = 3, [][]int{{0, 1}}, [][]int{{2, 1}}
	exp = []int{0, 1, -1}
	if r := shortestAlternatingPaths(n, redEdges, blueEdges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, redEdges, blueEdges = 1, [][]int{{0, 0}}, [][]int{{0, 0}}
	exp = []int{0}
	if r := shortestAlternatingPaths(n, redEdges, blueEdges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, redEdges, blueEdges = 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}
	exp = []int{0, 1, 2, 3, 7}
	if r := shortestAlternatingPaths(n, redEdges, blueEdges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
