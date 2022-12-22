package leetcode

import (
	"reflect"
	"testing"
)

func TestSumDistancesInTree(t *testing.T) {
	n, edges := 6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	exp := []int{8, 12, 6, 10, 10, 10}
	if r := sumOfDistancesInTree(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, edges = 1, [][]int{}
	exp = []int{0}
	if r := sumOfDistancesInTree(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, edges = 2, [][]int{{1, 0}}
	exp = []int{1, 1}
	if r := sumOfDistancesInTree(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}

	n, edges = 4, [][]int{{2, 0}, {3, 1}, {2, 1}}
	exp = []int{6, 4, 4, 6}
	if r := sumOfDistancesInTree(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
	n, edges = 5, [][]int{{0, 4}, {1, 3}, {1, 2}, {0, 2}}
	exp = []int{7, 7, 6, 10, 10}
	if r := sumOfDistancesInTree(n, edges); !reflect.DeepEqual(exp, r) {
		t.Fatalf("expect %v get %v", exp, r)
	}
}
