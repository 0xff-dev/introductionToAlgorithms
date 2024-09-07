package meituan

import "fmt"

/*
由n个节点构成，根结点为1的有根树，定义u和v相似：当且仅当u和v的子节点数量相同，输出相似节点对数。
*/
func problem2() {
	var (
		t    int
		node int
		a, b int
	)
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d", &node)
		adj := make(map[int][]int)
		for i := node; i > 1; i-- {
			fmt.Scanf("%d %d", &a, &b)
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
		v := make([]bool, node+1)
		q := []int{1}
		nodes := make(map[int]int)
		for len(q) > 0 {
			nq := make([]int, 0)
			for _, cur := range q {
				c := 0
				for _, rel := range adj[cur] {
					if v[rel] {
						continue
					}
					c++
					nq = append(nq, rel)
				}
				nodes[c]++
				v[cur] = true
			}
			q = nq
		}
		ans := 0
		for _, v := range nodes {
			if v > 1 {
				ans += (v - 1) * v / 2
			}
		}
		fmt.Println(ans)
	}
}
