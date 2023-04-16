package challengeprogrammingdatastructure

import "testing"

func TestDiameterOfTree(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1, 2},
		{1, 2, 1},
		{1, 3, 3},
	}
	if r := DiameterOfTree(n, edges); r != 5 {
		t.Fatalf("expect 5 get %d", r)
	}
}
