package challengeprogrammingdatastructure

func TopologicalSort(n int, edges [][]int) []int {
	nodes := make(map[int][]int)
	inDegree := make([]int, n)

	for _, e := range edges {
		from, to := e[0], e[1]
		inDegree[to]++
		if _, ok := nodes[from]; !ok {
			nodes[from] = make([]int, 0)
		}
		nodes[from] = append(nodes[from], to)
	}
	queue := []int{}
	ans := []int{}
	usedNodes := 0
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
			ans = append(ans, i)
			usedNodes++
		}
	}

	for len(queue) > 0 {
		nextQ := make([]int, 0)
		for _, item := range queue {
			for _, next := range nodes[item] {
				inDegree[next]--
				if inDegree[next] == 0 {
					nextQ = append(nextQ, next)
				}
			}
		}
		if len(nextQ) == 0 && usedNodes != n {
			// cycle
			return nil
		}
		queue = nextQ
		usedNodes += len(nextQ)
		ans = append(ans, nextQ...)
	}
	return ans
}
