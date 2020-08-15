package offer

import "testing"

func TestFindNumAppearOnce(t *testing.T) {
	input := []int{1,1,3,3,6,6,-1,-1, 8}
	if res := FindNumAppearOnce(input); res != 8 {
		t.Fatalf("expect result is %d get %d", 8, res)
	}
}

func TestTwoDiffNum(t *testing.T) {
	input := []int{1,1,2,3,3,5}
	if n1, n2 := TwoDiffNum(input); n1 != 5 && n2 != 2 {
		t.Fatalf("expect %d %d get %d %d", 2, 5, n1, n2)
	}
}
