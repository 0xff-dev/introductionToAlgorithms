package challengeprogrammingdatastructure

import "testing"

func TestCountingSort(t *testing.T) {
	nums := []int{2, 5, 1, 3, 2, 3, 0}
	target := CountingSort(nums)
	t.Logf("%v\n", target)
}
