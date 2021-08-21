package leetcode

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool {
	return true
}

func (this NestedInteger) GetInteger() int {
	return 0
}

func (this NestedInteger) SetInteger(value int) {

}
func (this NestedInteger) Add(ele NestedInteger) {

}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	Raw []*NestedInteger
	idx int
}

func Constructor3(nestedList []*NestedInteger) *NestedIterator {
	raw := make([]*NestedInteger, 0)
	deepTraverse(nestedList, &raw)
	return &NestedIterator{
		Raw: raw,
		idx: 0,
	}
}

func deepTraverse(nestedList []*NestedInteger, raw *[]*NestedInteger) {
	for idx, ele := range nestedList {
		if ele.IsInteger() {
			*raw = append(*raw, nestedList[idx])
			continue
		}
		deepTraverse(ele.GetList(), raw)
	}
}

func (this *NestedIterator) Next() int {
	r := this.Raw[this.idx].GetInteger()
	this.idx++
	return r
}

func (this *NestedIterator) HasNext() bool {
	return this.idx < len(this.Raw)
}
