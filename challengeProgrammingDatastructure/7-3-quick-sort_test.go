package challengeprogrammingdatastructure

import "testing"

func TestQuickSort(t *testing.T) {
	nums := []card{
		{
			'D', 3, 0,
		},
		{
			'H', 2, 1,
		},
		{
			'D', 1, 2,
		},
		{
			'S', 3, 3,
		},
		{
			'D', 2, 4,
		},
		{
			'C', 1, 5,
		},
	}

	QuickSort(nums, 0, 5)
	for _, item := range nums {
		t.Logf("%s\n", item)
	}
}
