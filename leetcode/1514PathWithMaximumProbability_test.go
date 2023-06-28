package leetcode

import "testing"

func TestMaxProbability(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1}, {1, 2}, {0, 2},
	}
	succProb := []float64{0.5, 0.5, 0.2}
	start, end := 0, 2
	if r := maxProbability(n, edges, succProb, start, end); r != 0.25000 {
		t.Fatalf("expect 0.25000 get %f", r)
	}

	succProb = []float64{0.5, 0.5, 0.3}
	if r := maxProbability(n, edges, succProb, start, end); r != 0.30000 {
		t.Fatalf("expect 0.30000 get %f", r)
	}

	edges = [][]int{
		{0, 1},
	}
	succProb = []float64{0.5}
	start, end = 0, 2
	if r := maxProbability(n, edges, succProb, start, end); r != 0.00000 {
		t.Fatalf("expect 0.00000 get %f", r)
	}
}
