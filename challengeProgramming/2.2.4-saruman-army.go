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

/*
	将一个模板切割成N块，每次切割的长度需要花销长度的能量，例如长度为21的要切割为5， 8， 8，切割为13和8的时候花销为21
	13切割为5，8的花销为13，所以总的花销为34

	反向思维，去最小的合成最终的长度。可以利用优先队列完成，小根堆
*/

func TotalCount(boards []int) int {
	length := len(boards)
	result := 0
	for length > 1 {
		l, r := 0, 1
		if boards[l] > boards[r] {
			l, r = r, l
		}
		for index := 2; index < length; index++ {
			if boards[l] > boards[index] {
				r, l = l, index
			} else if boards[r] > boards[index] {
				r = index
			}
		}
		tmp := boards[l] + boards[r]
		result += tmp
		if l == length-1 {
			l, r = r, l
		}
		boards[l], boards[r] = tmp, boards[length-1]
		length--
	}
	return result
}
