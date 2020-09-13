package challengeProgramming

import "testing"

func TestMostInvolvedWork(t *testing.T) {
	start, end := []int{1, 2, 4, 6, 8}, []int{3, 5, 7, 9, 10}
	if res := MostInvolvedWork(start, end); res != 3 {
		t.Fatalf("expect 3 get %d", res)
	}
}
