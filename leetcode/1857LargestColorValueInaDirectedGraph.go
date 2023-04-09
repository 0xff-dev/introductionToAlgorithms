package leetcode

type node1857 struct {
	in, out int
}

// 我的思路就是，先根据输入判断是否有环，
// 然后找到每个树，做dfs,自底向上, 超时
// 看了一个讲解是直接利用bfs拓扑直接判断了环，答案, 太牛了, 学习
func largestPathValue(colors string, edges [][]int) int {
	ans := -1
	nodes := len(colors)
	if len(edges) == 0 {
		return 1
	}
	nodeDegree := make([]node1857, nodes)
	nodeMap := make(map[int][]int)
	for _, e := range edges {
		nodeDegree[e[0]].out++
		nodeDegree[e[1]].in++
		if _, ok := nodeMap[e[0]]; !ok {
			nodeMap[e[0]] = make([]int, 0)
		}
		nodeMap[e[0]] = append(nodeMap[e[0]], e[1])
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		nd := make([]node1857, nodes)
		copy(nd, nodeDegree)
		queue := make([]int, 0)
		usedNodes := 0
		nodeValues := make([]int, nodes)
		for n, degree := range nd {
			if degree.in == 0 {
				queue = append(queue, n)
				usedNodes++
				if colors[n] == byte(ch) {
					nodeValues[n]++
				}
			}
		}
		for len(queue) > 0 {
			nextQ := make([]int, 0)
			for _, item := range queue {
				for _, to := range nodeMap[item] {
					nd[to].in--
					addOne := 0
					if colors[to] == byte(ch) {
						addOne = 1
					}
					if r := nodeValues[item] + addOne; r > nodeValues[to] {
						nodeValues[to] = r
					}
					if nodeValues[to] > ans {
						ans = nodeValues[to]
					}
					if nd[to].in == 0 {
						nextQ = append(nextQ, to)
					}
				}
			}
			if len(nextQ) == 0 && usedNodes != nodes {
				return -1
			}
			queue = nextQ
			usedNodes += len(nextQ)
		}
	}

	// edgeMap := make(map[int][]int)
	// target := make([]int, nodes)
	// for i := 0; i < nodes; i++ {
	// 	target[i] = -1
	// }
	// // paths := make([][]int, nodes)
	// // for i := 0; i < nodes; i++ {
	// // 	paths[i] = make([]int, 26)
	// // }
	// for _, edge := range edges {
	// 	if edge[0] == edge[1] {
	// 		return -1
	// 	}
	// 	target[edge[1]] = 0
	// 	if _, ok := edgeMap[edge[0]]; !ok {
	// 		edgeMap[edge[0]] = make([]int, 0)
	// 	}
	// 	edgeMap[edge[0]] = append(edgeMap[edge[0]], edge[1])
	// }
	// // 加一个cache, 嘿嘿嘿
	// var dfs func(int, map[int][26]int) [26]int
	// dfs = func(node int, cache map[int][26]int) [26]int {
	// 	next, ok := edgeMap[node]
	// 	if !ok {
	// 		ans := [26]int{}
	// 		ans[colors[node]-'a']++
	// 		return ans
	// 	}
	// 	if v, ok := cache[node]; ok {
	// 		return v
	// 	}
	// 	a := [26]int{}
	// 	for _, n := range next {
	// 		b := dfs(n, cache)
	// 		for i := 0; i < 26; i++ {
	// 			if b[i] > a[i] {
	// 				a[i] = b[i]
	// 			}
	// 		}
	// 	}
	// 	a[colors[node]-'a']++
	// 	cache[node] = a
	// 	return a
	// }
	// for node := 0; node < nodes; node++ {
	// 	if target[node] == -1 {
	// 		c := make(map[int][26]int)
	// 		a26 := dfs(node, c)
	// 		for _, i := range a26 {
	// 			if i > ans {
	// 				ans = i
	// 			}
	// 		}
	// 	}
	// }

	// 这个思路是直接考察每一条边加入进去以后，回产生的影响。
	// for _, edge := range edges {
	// 	from, to := edge[0], edge[1]
	// 	fromColr := colors[from] - 'a'
	// 	toColor := colors[to] - 'a'
	// 	if paths[from][fromColr] == 0 {
	// 		paths[from][fromColr] = 1
	// 	}
	// 	if r := paths[from][toColor] + 1; r > paths[to][toColor] {
	// 		paths[to][toColor] = r
	// 	}
	// 	if paths[to][toColor] > ans {
	// 		ans = paths[to][toColor]
	// 	}
	// 	for i := 0; i < 26; i++ {
	// 		if i == int(toColor) {
	// 			continue
	// 		}
	// 		if paths[from][i] > paths[to][i] {
	// 			paths[to][i] = paths[from][i]
	// 		}
	// 		if paths[to][i] > ans {
	// 			ans = paths[to][i]
	// 		}
	// 	}
	// }

	return ans
}
