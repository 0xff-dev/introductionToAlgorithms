package challengeprogrammingdatastructure

import "fmt"

// 直接遍历数组即可
func PrintComplteBinaryTree(nodes []int) {
	nodeIndex := 1
	queue := []int{0}
	for len(queue) > 0 {
		nextQ := make([]int, 0)
		for _, item := range queue {
			fmt.Printf("node %d: key = %d", nodeIndex, nodes[item])
			nodeIndex++
			if item != 0 {
				p := item / 2
				if item&1 == 0 {
					p--
				}
				fmt.Printf(", parent key = %d", nodes[p])
			}
			if left := item*2 + 1; left < len(nodes) {
				fmt.Printf(", left key = %d", nodes[left])
				nextQ = append(nextQ, left)
			}
			if right := item*2 + 2; right < len(nodes) {
				fmt.Printf(", right key = %d ", nodes[right])
				nextQ = append(nextQ, right)
			}
			fmt.Println()
		}
		queue = nextQ
	}
}
