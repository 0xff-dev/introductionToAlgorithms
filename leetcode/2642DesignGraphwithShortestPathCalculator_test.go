package leetcode

import "testing"

func TestConstructor2642(t *testing.T) {
	c := Constructor2642(4, [][]int{
		{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3},
	})
	if r := c.ShortestPath(3, 2); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
	if r := c.ShortestPath(0, 3); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}
	c.AddEdge([]int{1, 3, 4})
	if r := c.ShortestPath(0, 3); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}
}
