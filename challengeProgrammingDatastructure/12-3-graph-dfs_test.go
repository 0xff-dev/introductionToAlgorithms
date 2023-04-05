package challengeprogrammingdatastructure

import (
	"testing"
)

func TestGraphDFS(t *testing.T) {
	table := []adjTable{
		{1, 2, []int{2, 3}},
		{2, 2, []int{3, 4}},
		{3, 1, []int{5}},
		{4, 1, []int{6}},
		{5, 1, []int{6}},
		{6, 0, nil},
	}

	GraphDFS(table)
	table = []adjTable{
		{1, 2, []int{2, 3}},
		{2, 0, nil},
		{3, 0, nil},
	}
	GraphDFS(table)
}
