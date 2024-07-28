package leetcode

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	q := make([][2]int, 0)
	dist1 := make([]int, n+1)
	dist2 := make([]int, n+1)
	for i := range dist1 {
		dist1[i] = -1
		dist2[i] = -1
	}

	q = append(q, [2]int{1, 1})
	dist1[1] = 0

	for len(q) > 0 {
		node, freq := q[0][0], q[0][1]
		q = q[1:]

		timeTaken := dist1[node]
		if freq == 2 {
			timeTaken = dist2[node]
		}

		if (timeTaken/change)%2 == 1 {
			timeTaken = change*((timeTaken/change)+1) + time
		} else {
			timeTaken = timeTaken + time
		}

		for _, neighbor := range adj[node] {
			if dist1[neighbor] == -1 {
				dist1[neighbor] = timeTaken
				q = append(q, [2]int{neighbor, 1})
			} else if dist2[neighbor] == -1 && dist1[neighbor] != timeTaken {
				if neighbor == n {
					return timeTaken
				}
				dist2[neighbor] = timeTaken
				q = append(q, [2]int{neighbor, 2})
			}
		}
	}
	return 0
}
