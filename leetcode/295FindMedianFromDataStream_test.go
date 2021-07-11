package leetcode

import "testing"

type mediaStruct struct {
	Nums []int
	Ans  float64
}

func TestConstructor1(t *testing.T) {
	t.Log("odd test...")
	for _, ts := range []mediaStruct{
		{[]int{1, 2, 3}, 2},
		{[]int{-4, -3, -2, -1, 0, 1, 2}, -1},
		{[]int{-3, -2, -1}, -2},
		{[]int{1}, 1},
		{[]int{-2, -2, -1, -1, 0}, -1},
	} {
		a := Constructor1()
		for _, n := range ts.Nums {
			a.AddNum(n)
		}
		if r := a.FindMedian(); r != ts.Ans {
			t.Fatalf("input: %v expect %f get %f", ts.Nums, ts.Ans, r)
		}
	}

	t.Log("even test...")
	for _, ts := range []mediaStruct{
		{[]int{1, 2}, 1.5},
		{[]int{-2, -2, -2, 1, 1, 3}, -0.5},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{-2, -2, 1, 1, 2, 2}, 1.0},
		{[]int{-1000, -100, -100, 0, 2, 2, 3, 4}, 1.0},
	} {
		a := Constructor1()
		for _, n := range ts.Nums {
			a.AddNum(n)
		}
		if r := a.FindMedian(); r != ts.Ans {
			t.Fatalf("input: %v expect %f get %f", ts.Nums, ts.Ans, r)
		}
	}
}
