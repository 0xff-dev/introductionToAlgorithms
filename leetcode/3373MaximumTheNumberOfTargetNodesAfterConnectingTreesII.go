package leetcode

func maxTargetNodes3373(edges1 [][]int, edges2 [][]int) []int {
	var dfs func(node, parent, depth int, children [][]int, color []int) int
	dfs = func(node, parent, depth int, children [][]int, color []int) int {
		res := 1 - depth%2
		color[node] = depth % 2
		for _, child := range children[node] {
			if child == parent {
				continue
			}
			res += dfs(child, node, depth+1, children, color)
		}
		return res
	}

	build := func(edges [][]int, color []int) []int {
		n := len(edges) + 1
		children := make([][]int, n)
		for _, edge := range edges {
			u, v := edge[0], edge[1]
			children[u] = append(children[u], v)
			children[v] = append(children[v], u)
		}
		res := dfs(0, -1, 0, children, color)
		return []int{res, n - res}
	}

	n := len(edges1) + 1
	m := len(edges2) + 1
	color1 := make([]int, n)
	color2 := make([]int, m)
	count1 := build(edges1, color1)
	count2 := build(edges2, color2)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = count1[color1[i]] + max(count2[0], count2[1])
	}
	return res
}
