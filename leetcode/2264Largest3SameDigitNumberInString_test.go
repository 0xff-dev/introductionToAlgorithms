package leetcode

import "testing"

func TestLargestGoodInteger(t *testing.T) {
	if r := largestGoodInteger("6777133339"); r != "777" {
		t.Fatalf("expect 777 get %s", r)
	}
	if r := largestGoodInteger("2300019"); r != "000" {
		t.Fatalf("expect 000 get %s", r)
	}

	if r := largestGoodInteger("42352338"); r != "" {
		t.Fatalf("expect '' get %s", r)
	}
}
