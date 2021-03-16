package leetcode

type sampleNode struct {
	Val   int
	Left  *sampleNode
	Right *sampleNode
	Next  *sampleNode
}

func connect(root *sampleNode) *sampleNode {
	if root == nil {
		return root
	}

	level := []*sampleNode{root}
	for len(level) > 0 {
		tmp := make([]*sampleNode, 0)
		for idx := 0; idx < len(level); idx++ {
			if idx+1 < len(level) {
				level[idx].Next = level[idx+1]
			}
			if level[idx].Left != nil {
				tmp = append(tmp, level[idx].Left)
			}
			if level[idx].Right != nil {
				tmp = append(tmp, level[idx].Right)
			}
		}
		level = tmp

	}
	return root
}
