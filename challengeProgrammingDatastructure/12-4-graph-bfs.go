package challengeprogrammingdatastructure

import "fmt"

func GraphBFS(table []adjTable) {
	distance := 1
	queue := []int{1}
	visited := make([]int, len(table))
	for len(queue) > 0 {
		nextQ := make([]int, 0)
		for _, item := range queue {
			if visited[item-1] == 0 {
				visited[item-1] = distance
			}

			for _, c := range table[item-1].points {
				if visited[c-1] != 0 {
					continue
				}
				nextQ = append(nextQ, c)
			}
		}
		distance++
		queue = nextQ
	}

	for i := 0; i < len(table); i++ {
		fmt.Printf("%d %d\n", i+1, visited[i]-1)
	}
}
