package challengeprogrammingdatastructure

import (
	"container/heap"
)

// 树是没有环的图
// 图G的生成树(SpanningTree) G` 是图G的子图，它拥有图G所有的顶点且在保证自身为树的前提下拥有尽可能多的边
// 最短路径问题分为 - 单源最短路径和全点对最短路径
func MinimumSpanningTree(n int, graph [][]int) (int, int) {
	ans := Prim(n, graph)
	ans2 := Kruskal(n, graph)
	return ans, ans2
}

func Prim(n int, graph [][]int) int {
	visited := make([]bool, n)
	ans := 0
	visited[0] = true
	for {
		distance := -1
		selectNode := -1
		for i, v := range visited {
			if !v {
				continue
			}
			for idx, rel := range graph[i] {
				if visited[idx] || rel == -1 {
					continue
				}
				if distance == -1 || distance > rel {
					distance = rel
					selectNode = idx
				}
			}
		}
		if distance == -1 {
			break
		}
		ans += distance
		visited[selectNode] = true
	}
	return ans
}

type edge struct {
	s, e, w int
}

type edges []edge

func (es *edges) Len() int {
	return len(*es)
}

func (es *edges) Less(i, j int) bool {
	return (*es)[i].w < (*es)[j].w
}

func (es *edges) Swap(i, j int) {
	(*es)[i], (*es)[j] = (*es)[j], (*es)[i]
}

func (es *edges) Push(x interface{}) {
	*es = append(*es, x.(edge))
}

func (es *edges) Pop() interface{} {
	old := *es
	l := len(old)
	x := old[l-1]
	*es = old[:l-1]
	return x
}

func Kruskal(n int, graph [][]int) int {
	u := uinionFind125{father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}
	es := edges([]edge{})
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] != -1 {
				heap.Push(&es, edge{i, j, graph[i][j]})
			}
		}
	}

	usedNodes := make([]bool, n)
	for notUseAllNodes(usedNodes) {
		x := heap.Pop(&es).(edge)
		fs := u.findFather(x.s)
		fe := u.findFather(x.e)
		if fs != fe {
			usedNodes[x.s] = true
			usedNodes[x.e] = true
			u.union(fs, fe)
			ans += x.w
		}
	}
	return ans
}

func notUseAllNodes(nodes []bool) bool {
	for _, v := range nodes {
		if !v {
			return true
		}
	}
	return false
}
