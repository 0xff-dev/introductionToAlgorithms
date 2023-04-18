package leetcode

func intersect(quadTree1 *QuadNode, quadTree2 *QuadNode) *QuadNode {
	if quadTree1.IsLeaf && quadTree2.IsLeaf {
		return &QuadNode{IsLeaf: true, Val: quadTree1.Val || quadTree2.Val}
	}
	node := &QuadNode{IsLeaf: quadTree1.IsLeaf || quadTree2.IsLeaf}
	if quadTree1.IsLeaf {
		if !quadTree1.Val {
			return quadTree2
		}
		node.Val = true
		return node
	}
	if quadTree2.IsLeaf {
		if !quadTree2.Val {
			return quadTree1
		}
		node.Val = true
		return node
	}
	bl := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	br := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	tl := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	tr := intersect(quadTree1.TopRight, quadTree2.TopRight)
	if bl.IsLeaf && br.IsLeaf && tl.IsLeaf && tr.IsLeaf {
		if bl.Val == br.Val && bl.Val == tl.Val && bl.Val == tr.Val {
			node.IsLeaf = true
			node.Val = bl.Val
			return node
		}
	}
	node.BottomLeft = bl
	node.BottomRight = br
	node.TopLeft = tl
	node.TopRight = tr
	return node
}
