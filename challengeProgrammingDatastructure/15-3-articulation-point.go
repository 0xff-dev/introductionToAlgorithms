package challengeprogrammingdatastructure

import "fmt"

func dfs153(edges map[int][]int, current, prev int, timer *int, prenum, parent, lowest []int, visited []bool) {
	prenum[current], lowest[current] = *timer, *timer
	*timer++

	visited[current] = true
	for _, next := range edges[current] {
		if !visited[next] {
			parent[next] = current
			dfs153(edges, next, current, timer, prenum, parent, lowest, visited)
		} else if next != prev {
			if prenum[next] < lowest[current] {
				lowest[current] = prenum[next]
			}
		}
	}
	return
}

func ArticulationPoint(n int, edges [][]int) {
	prenum := make([]int, n)
	parent := make([]int, n)
	lowest := make([]int, n)
	visited := make([]bool, n)

	timer := 1
	g := make(map[int][]int)
	for _, e := range edges {
		s, t := e[0], e[1]
		if _, ok := g[s]; !ok {
			g[s] = make([]int, 0)
		}
		if _, ok := g[t]; !ok {
			g[t] = make([]int, 0)
		}
		g[s] = append(g[s], t)
		g[t] = append(g[t], s)
	}

	dfs153(g, 0, -1, &timer, prenum, parent, lowest, visited)
	ap := make(map[int]struct{})
	np := 0
	for i := 1; i < n; i++ {
		p := parent[i]
		if p == 0 {
			np++
			continue
		}
		if lowest[i] >= prenum[p] {
			ap[p] = struct{}{}
		}
	}

	if np > 1 {
		ap[0] = struct{}{}
	}
	for key := range ap {
		fmt.Println(key)
	}
}
