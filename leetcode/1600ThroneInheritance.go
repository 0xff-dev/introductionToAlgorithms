package leetcode

type kingTree struct {
	name  string
	live  bool
	child []*kingTree
}

func (k *kingTree) preOrder(out *[]string) {
	if k == nil {
		return
	}
	if k.live {
		*out = append(*out, k.name)
	}
	for _, c := range k.child {
		c.preOrder(out)
	}
}

type ThroneInheritance struct {
	root      *kingTree
	nodeCache map[string]*kingTree
}

func Constructor1600(kingName string) ThroneInheritance {
	king := &kingTree{
		name:  kingName,
		live:  true,
		child: make([]*kingTree, 0),
	}

	cache := map[string]*kingTree{
		kingName: king,
	}
	ti := ThroneInheritance{
		root:      king,
		nodeCache: cache,
	}
	return ti
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	node := this.nodeCache[parentName]
	childNode := &kingTree{
		name:  childName,
		live:  true,
		child: make([]*kingTree, 0),
	}
	node.child = append(node.child, childNode)
	this.nodeCache[childName] = childNode

}

func (this *ThroneInheritance) Death(name string) {
	this.nodeCache[name].live = false
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var out []string
	this.root.preOrder(&out)
	return out
}
