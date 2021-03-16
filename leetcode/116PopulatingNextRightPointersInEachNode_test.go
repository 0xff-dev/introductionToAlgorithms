package leetcode

import "testing"

func TestConnect(t *testing.T) {
	tree := &sampleNode{1, &sampleNode{2, nil, nil, nil}, &sampleNode{3, nil, nil, nil}, nil}
	tree = connect(tree)

}
