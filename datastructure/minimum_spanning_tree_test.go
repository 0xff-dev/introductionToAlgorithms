package datastructure

import (
	"fmt"
	"testing"
)

func tree() []edge {
	edges := make([]edge, 0)
	edges = append(edges, edge{'a', 'b', 4})
	edges = append(edges, edge{'a', 'h', 8})
	edges = append(edges, edge{'b', 'h', 11})
	edges = append(edges, edge{'c', 'b', 8})
	edges = append(edges, edge{'c', 'i', 2})
	edges = append(edges, edge{'i', 'h', 7})
	edges = append(edges, edge{'h', 'g', 1})
	edges = append(edges, edge{'i', 'g', 6})
	edges = append(edges, edge{'f', 'c', 4})
	edges = append(edges, edge{'g', 'f', 2})
	edges = append(edges, edge{'c', 'd', 7})
	edges = append(edges, edge{'f', 'd', 14})
	edges = append(edges, edge{'e', 'd', 9})
	edges = append(edges, edge{'e', 'f', 10})
	return edges
}

func TestKruskal(t *testing.T) {
	edges := tree()
	e, w := kruskal(edges)
	t.Log("weight: ", w)
	for _, item := range e {
		fmt.Print(item)
	}
}

func TestPrim(t *testing.T) {
	edges := tree()
	visited, r := prim(edges, 9)
	t.Log(r)
	for e := range visited {
		fmt.Print(string(e))
	}
}
