package offer

import "testing"

func TestIsPopOrder(t *testing.T) {
	order := []int{4,5,3,2,1}
	in := []int{1,2,3,4,5}
	if IsPopOrder(order, in) != true {
		t.Fatalf("except true")
	}
	order = []int{4,3,5,1,2}
	if IsPopOrder(order, in) != false {
		t.Fatalf("except false")
	}
}
