package challengeprogrammingdatastructure

import "testing"

func TestGraphAdjMatrix(t *testing.T) {
	table := []adjTable{
		{1, 2, []int{2, 4}},
		{2, 1, []int{4}},
		{3, 0, nil},
		{4, 1, []int{3}},
	}
	GraphAdjMatrix(4, table)
}
