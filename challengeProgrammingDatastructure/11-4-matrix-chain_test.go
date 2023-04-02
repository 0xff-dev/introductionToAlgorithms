package challengeprogrammingdatastructure

import "testing"

func TestMatrixChainMulti(t *testing.T) {
	m := []Matrix{
		{30, 35}, {35, 15}, {15, 5}, {5, 10}, {10, 20}, {20, 25},
	}
	if r := MatrixChainMulti(m); r != 15125 {
		t.Fatalf("expect 15125 get %d", r)
	}

	m = []Matrix{
		{1, 2}, {2, 3},
	}
	if r := MatrixChainMulti(m); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
