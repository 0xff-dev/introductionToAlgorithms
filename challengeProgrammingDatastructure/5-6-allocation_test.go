package challengeprogrammingdatastructure

import "testing"

func TestGoodsAlloc(t *testing.T) {
	k := 3
	weight := []int{8, 1, 7, 3, 9}
	if r := goodsAlloc(k, weight); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
	k = 3
	weight = []int{12, 8, 8}
	if r := goodsAlloc(k, weight); r != 12 {
		t.Fatalf("expect 12 get %d", r)
	}
}
