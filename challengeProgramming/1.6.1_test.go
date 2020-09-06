package challengeProgramming

import "testing"

type testStruct struct {
	Woods  []int
	Result int
}

func TestTrianglePerimeter(t *testing.T) {
	for _, testItem := range []testStruct{
		{
			[]int{2, 3, 4, 5, 10},
			12,
		},
		{
			[]int{4, 5, 10, 20},
			0,
		},
	} {
		if res := TrianglePerimeter(testItem.Woods); res != testItem.Result {
			t.Fatalf("expect %d get %d", testItem.Result, res)
		}
	}
}
