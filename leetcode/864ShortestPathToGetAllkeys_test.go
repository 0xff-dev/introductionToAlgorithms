package leetcode

import (
	"testing"
)

func TestShortestPathAllKeys(t *testing.T) {
	grid := []string{
		"@.a..", "###.#", "b.A.B",
	}
	if r := shortestPathAllKeys(grid); r != 8 {
		t.Fatalf("expect 8 get %d", r)
	}

	grid = []string{"@..aA", "..B#.", "....b"}
	if r := shortestPathAllKeys(grid); r != 6 {
		t.Fatalf("expect 6 get %d", r)
	}

	grid = []string{"@Aa"}
	if r := shortestPathAllKeys(grid); r != -1 {
		t.Fatalf("expect -1 get %d", r)
	}

	grid = []string{"@...a", ".###A", "b.BCc"}
	if r := shortestPathAllKeys(grid); r != 10 {
		t.Fatalf("expect 10 get %d", r)
	}
}
