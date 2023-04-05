package challengeprogrammingdatastructure

import (
	"fmt"
	"sort"
)

func GraphAdjMatrix(n int, table []adjTable) {
	sort.Slice(table, func(i, j int) bool {
		return table[i].u < table[j].u
	})

	for _, row := range table {
		sort.Ints(row.points)
		index := 0
		for i := 0; i < n; i++ {
			if index < len(row.points) && i+1 == row.points[index] {
				fmt.Print("1 ")
				index++
				continue
			}
			fmt.Print("0 ")
		}
		fmt.Println()
	}
}
