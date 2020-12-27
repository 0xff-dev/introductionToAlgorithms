package datastructure

import (
	"fmt"
	"sort"
)

type forest struct {
	F []map[byte]struct{}
}

func (f *forest) FindSet(a byte) int {
	for idx := 0; idx < len(f.F); idx++ {
		if _, ok := f.F[idx][a]; ok {
			return idx
		}
	}
	return -1
}

func (f *forest) Union(a, b byte) {
	aIdx, bIdx := f.FindSet(a), f.FindSet(b)
	if aIdx != bIdx {
		// union
		for k := range f.F[bIdx] {
			f.F[aIdx][k] = struct{}{}
		}

		tmpList := make([]map[byte]struct{}, 0)
		for i := 0; i < len(f.F); i++ {
			if i == bIdx {
				continue
			}
			tmpList = append(tmpList, f.F[i])
		}
		f.F = tmpList
	}
}

func (f *forest) AddTree(a byte) {
	f.F = append(f.F, map[byte]struct{}{
		a: {},
	})
}

type edge struct {
	s, e byte
	w    int
}

func (e edge) String() string {
	return fmt.Sprintf("edge: %c -- edge: %c -- weight: %d\n", e.s, e.e, e.w)
}

func initForest(edges []edge) *forest {
	// first ,sort edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	f := &forest{F: []map[byte]struct{}{}}
	for _, e := range edges {
		if f.FindSet(e.s) == -1 {
			f.AddTree(e.s)
		}
		if f.FindSet(e.e) == -1 {
			f.AddTree(e.e)
		}
	}
	return f
}

// can use priority queue
func adjacentNodes(edges []edge, visited map[byte]struct{}) map[byte]int {
	r := make(map[byte]int)
	for e := range visited {
		for _, item := range edges {
			if item.s == e {
				if _, ok := visited[item.e]; !ok {
					v, ok := r[item.e]
					if !ok || v > item.w {
						r[item.e] = item.w
					}
				}
			}
			if item.e == e {
				if _, ok := visited[item.s]; !ok {
					v, ok := r[item.s]
					if !ok || v > item.w {
						r[item.s] = item.w
					}
				}
			}
		}
	}
	return r
}

func minDist(nodes map[byte]int) (byte, int) {
	min, minNode := INF, byte(' ')
	for k, v := range nodes {
		if v < min {
			min, minNode = v, k
		}
	}
	return minNode, min
}

// kruskal algorithm
func kruskal(edges []edge) ([]edge, int) {
	f := initForest(edges)
	weight := 0
	newEdges := make([]edge, 0)
	for _, e := range edges {
		sIdx, eIdx := f.FindSet(e.s), f.FindSet(e.e)
		if sIdx != eIdx {
			// union
			newEdges = append(newEdges, e)
			weight += e.w
			f.Union(e.s, e.e)
		}
	}
	return newEdges, weight
}

// prim algorithm
func prim(edges []edge, nodeCount int) (map[byte]struct{}, int) {
	visited := make(map[byte]struct{})
	visited[edges[0].s] = struct{}{}
	count := 1
	r := 0
	for count < nodeCount {
		nodes := adjacentNodes(edges, visited)
		nodeByte, dist := minDist(nodes)
		r += dist
		visited[nodeByte] = struct{}{}
		count++
	}
	return visited, r
}
