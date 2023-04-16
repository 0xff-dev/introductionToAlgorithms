package challengeprogrammingdatastructure

import "testing"

func TestArticulationPoint(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2}, {2, 3},
	}
	ArticulationPoint(n, edges)
}
