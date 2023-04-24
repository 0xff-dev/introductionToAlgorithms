package challengeprogrammingdatastructure

import "testing"

func TestPrimeNumbers(t *testing.T) {
	nums := []int{2, 3, 4, 5, 6, 7}
	if r := PrimeNumbers(nums); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}
}
