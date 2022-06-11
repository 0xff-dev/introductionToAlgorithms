package leetcode

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	//nums := []int{2, 3, 1, 2, 4, 3}
	//if r := minSubArrayLen(7, nums); r != 2 {
	//	t.Fatalf("expect 2 get %d", r)
	//}
	//
	//nums = []int{8, 3}
	//if r := minSubArrayLen(3, nums); r != 1 {
	//	t.Fatalf("expect 1 get %d", r)
	//}
	//
	//nums = []int{1, 4, 4}
	//if r := minSubArrayLen(4, nums); r != 1 {
	//	t.Fatalf("expect 1 get %d", r)
	//}
	//
	//nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	//if r := minSubArrayLen(11, nums); r != 0 {
	//	t.Fatalf("expect 0 get %d", r)
	//}

	nums := []int{1, 2, 3, 4, 5}
	if r := minSubArrayLen(11, nums); r != 3 {
		t.Fatalf("expect 3 get %d", r)
	}
}
