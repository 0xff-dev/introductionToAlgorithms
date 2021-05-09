package leetcode

import "testing"

func TestMinHeightShelves(t *testing.T) {
	books := [][]int{
		{1, 1},
		{2, 3},
		{2, 3},
	}
	if r := minHeightShelves(books, 4); r != 4 {
		t.Fatalf("expect 4 get %d", r)
	}

	books = [][]int{
		{1, 1},
		{2, 3},
	}

	if r := minHeightShelves(books, 4); r != 3 {
		t.Fatalf("expect 3 get %d", 4)
	}

	books = [][]int{
		{1, 1},
		{2, 3},
		{2, 3},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 2},
	}

	if r := minHeightShelves(books, 4); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	books = [][]int{
		{1, 2},
	}

	if r := minHeightShelves(books, 1); r != 2 {
		t.Fatalf("expect 2 get %d", r)
	}

	books = [][]int{
		{1, 1},
		{2, 3},
		{2, 3},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 2},
		{3, 2},
	}

	if r := minHeightShelves(books, 5); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

}
