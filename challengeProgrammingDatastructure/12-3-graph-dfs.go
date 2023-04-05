package challengeprogrammingdatastructure

import "fmt"

func GraphDFS(table []adjTable) {
	s, e := make([]int, len(table)), make([]int, len(table))
	s[0] = 1
	graphDFS(1, s, e, table)
	for i := 0; i < len(table); i++ {
		fmt.Printf("%d %d %d\n", i+1, s[i], e[i])
	}
	fmt.Println()
}

func graphDFS(start int, s, e []int, table []adjTable) {
	if table[start-1].k == 0 {
		e[start-1] = s[start-1] + 1
		return
	}
	nextStart := s[start-1]
	for _, c := range table[start-1].points {
		if s[c-1] != 0 {
			continue
		}
		s[c-1] = nextStart + 1
		graphDFS(c, s, e, table)
		nextStart = e[c-1]
	}
	e[start-1] = nextStart + 1
}
