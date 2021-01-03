package datastructure

import "fmt"

func initSingleSource(edges []edge, start byte, path map[byte]byte) map[byte]int {
	r := make(map[byte]int)
	r[start] = 0
	for _, e := range edges {
		path[e.s], path[e.e] = e.s, e.e
		if e.s == start {
			r[e.e] = e.w
			continue
		}
		r[e.s], r[e.e] = INF, INF
	}
	return r
}

func relax(r map[byte]int, e edge, path map[byte]byte) {
	if r[e.e] == INF {
		path[e.e] = e.s
		if r[e.s] == INF {
			r[e.e] = e.w
			return
		}
		r[e.e] = r[e.s] + e.w
		return
	}
	if r[e.s] != INF && r[e.e] > r[e.s]+e.w {
		path[e.e] = e.s
		r[e.e] = r[e.s] + e.w
	}
}

func BellmanFord(edges []edge, s, e byte) (bool, int) {
	if len(edges) == 0 {
		return true, 0
	}
	path := make(map[byte]byte)
	r := initSingleSource(edges, s, path)
	for i := 1; i < len(edges); i++ {
		for _, e := range edges {
			relax(r, e, path)
		}
	}

	// todo Why is this judgment
	for _, e := range edges {
		if r[e.e] > r[e.s]+e.w {
			return false, -1
		}
	}
	findPath(s, e, path)

	return true, r[e]
}

func findPath(start, end byte, paths map[byte]byte) {
	r := []byte{end}
	for end != start {
		end = paths[end]
		r = append(r, end)
	}
	for idx := len(r) - 1; idx > 0; idx-- {
		fmt.Printf("%c->", r[idx])
	}
	fmt.Printf("%c", r[0])
}
