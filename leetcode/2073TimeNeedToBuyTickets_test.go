package leetcode

import "testing"

func TestTimeRequiredToBuy(t *testing.T) {
	tickets := []int{2, 3, 2}
	k := 2
	exp := 6
	if r := timeRequiredToBuy(tickets, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	tickets = []int{5, 1, 1, 1}
	k = 0
	exp = 8
	if r := timeRequiredToBuy(tickets, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	tickets = []int{84, 49, 5, 24, 70, 77, 87, 8}
	k = 3
	exp = 154
	if r := timeRequiredToBuy(tickets, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}

	tickets = []int{83, 86, 38, 31, 59, 25, 89, 71, 54, 71, 84}
	k = 1
	exp = 687
	if r := timeRequiredToBuy(tickets, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
	tickets = []int{62, 30, 37, 75, 7, 17, 33, 48, 41, 63, 54, 67, 14, 38, 47, 23, 22, 50, 50, 59, 67, 55, 17, 51, 46, 26}
	k = 21
	exp = 1036
	if r := timeRequiredToBuy(tickets, k); r != exp {
		t.Fatalf("expect %d get %d", exp, r)
	}
}
