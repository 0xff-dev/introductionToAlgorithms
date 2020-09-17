package challengeProgramming

import (
	"fmt"
	"sort"
)

/*
	直线上有N个点，点i的位置Xi，从N个点中选择若干添加标记，对于每个点保证其范围R内要有标记的点。自己算自我满足，希望尽量少的添加标记.
*/

// 1 7| 15 20, 30, 50
func MarkNode(nodes []int, r int) {
	if len(nodes) == 0 {
		return
	}
	sort.Ints(nodes)
	markNodes, start, flag := make([]int, 0), nodes[0], 0
	for index := 0; index < len(nodes); index++ {
		if nodes[index] <= start+r {
			continue
		}
		flag = 1 - flag
		if flag&1 == 0 {
			markNodes = append(markNodes, start)
			start = nodes[index]
		} else {
			start = nodes[index-1]
			index--
		}
	}
	if flag == 0 || len(markNodes) == 0 {
		markNodes = append(markNodes, start)
	}
	fmt.Println("mark nodes: ", markNodes)
}
