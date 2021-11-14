package leetcode

type CombinationIterator struct {
	String          string
	Yield           []string
	Len, Total, Now int
}

func ConstructorForCom(characters string, combinationLength int) CombinationIterator {
	o := CombinationIterator{
		String: characters,
		Yield:  make([]string, 0),
		Len:    combinationLength,
		Total:  0,
		Now:    0,
	}
	(&o).construct(0, 0, []byte{})
	return o
}

func (this *CombinationIterator) construct(length, now int, paths []byte) {
	if now > len(this.String) {
		return
	}
	if length == this.Len {
		this.Total++
		this.Yield = append(this.Yield, string(paths))
		return
	}
	for ch := now; ch < len(this.String); ch++ {
		this.construct(length+1, ch+1, append(paths, this.String[ch]))
	}
}

func (this *CombinationIterator) Next() string {
	if this.HasNext() {
		return_t := this.Now
		this.Now++
		return this.Yield[return_t]
	}
	return ""
}

func (this *CombinationIterator) HasNext() bool {
	return this.Now < this.Total
}
