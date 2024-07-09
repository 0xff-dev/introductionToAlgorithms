package leetcode

import "testing"

const (
	within = 1e-5
)

func TestAverageWaitingTime(t *testing.T) {
	customers := [][]int{
		{1, 2}, {2, 5}, {4, 3},
	}
	exp := float64(5.00000)
	r := averageWaitingTime(customers)
	diff := r - exp
	if diff < 0 {
		diff = -diff
	}
	if diff > within {
		t.Fatalf("expect %f get %f", exp, r)
	}

	customers = [][]int{
		{5, 2}, {5, 4}, {10, 3}, {20, 1},
	}
	exp = 3.25000
	r = averageWaitingTime(customers)
	diff = r - exp
	if diff < 0 {
		diff = -diff
	}
	if diff > within {
		t.Fatalf("expect %f get %f", exp, r)
	}
}
