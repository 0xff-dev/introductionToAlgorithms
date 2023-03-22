package challengeprogrammingdatastructure

import (
	"testing"
)

func TestExhaustiveSearch(t *testing.T) {
	nums := []int{1, 5, 7, 10, 21}
	var cases = []struct {
		target int
		exp    bool
	}{
		{2, false},
		{4, false},
		{17, true},
		{8, true},
		{44, true},
	}
	for _, c := range cases {
		if r := ExhaustiveSearch(nums, c.target); r != c.exp {
			t.Fatalf("expect %v get %v", c.exp, r)
		}
	}
}
