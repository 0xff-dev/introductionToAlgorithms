package challengeprogrammingdatastructure

func MinimumSpanningTreeEdges(n int, edges [][]int) int {
	visited := make([]bool, n)
	usedNodes := 0
	edgeMap := make(map[int][]int)
	for _, e := range edges {
		s, t, w := e[0], e[1], e[2]
		if _, ok := edgeMap[s]; !ok {
			edgeMap[s] = make([]int, 0)
		}
		if _, ok := edgeMap[t]; !ok {
			edgeMap[t] = make([]int, 0)
		}
		edgeMap[s] = append(edgeMap[s], t, w)
		edgeMap[t] = append(edgeMap[t], s, w)
	}

	visited[0] = true
	ans := 0
	for usedNodes < n-1 {
		dis, selectNode := -1, -1
		for node, v := range visited {
			if !v {
				continue
			}
			for i := 0; i < len(edgeMap[node]); i += 2 {
				t, w := edgeMap[node][i], edgeMap[node][i+1]
				if !visited[t] && (dis == -1 || dis > w) {
					selectNode = t
					dis = w
				}
			}
		}
		usedNodes++
		visited[selectNode] = true
		ans += dis
	}
	return ans
}
