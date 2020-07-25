package offer

import "testing"

func TestCounterclockwisePrint(t *testing.T) {
	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	CounterclockwisePrint(array)

	array = [][]int{
		{1, 2, 3, 4},
	}
	CounterclockwisePrint(array)

	array = [][]int{
		{1},
		{2},
		{3},
		{4},
	}
	CounterclockwisePrint(array)

	array = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	CounterclockwisePrint(array)

	CounterclockwisePrint([][]int{})

	array = [][]int{
		{1},
	}

	CounterclockwisePrint(array)
}
