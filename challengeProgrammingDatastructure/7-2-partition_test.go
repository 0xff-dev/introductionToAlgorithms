package challengeprogrammingdatastructure

import "testing"

func TestPartition(t *testing.T) {
	nums := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	r := Partition(nums, 0, 11)
	t.Logf("%d\n", r)

	nums = []int{4, 3, 2, 1}
	r = Partition(nums, 0, 3)
	t.Logf("%d\n", r)
	nums = []int{1, 2, 3, 4}
	r = Partition(nums, 0, 3)
	t.Logf("%d\n", r)
}
