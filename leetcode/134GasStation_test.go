package leetcode

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	if r := canCompleteCircuit(gas, cost); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}

	gas, cost = []int{2, 3, 4}, []int{3, 4, 3}
	if r := canCompleteCircuit(gas, cost); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

}
