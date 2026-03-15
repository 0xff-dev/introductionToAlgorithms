package leetcode

type Fancy struct {
	source   []int
	add, mul int64

	mod int64
}

func Constructor1622() Fancy {
	return Fancy{
		source: []int{},
		add:    0,
		mul:    1,
		mod:    1000000007,
	}
}

func (this *Fancy) pow(a, b int64) int64 {
	res := int64(1)
	a %= this.mod
	for b > 0 {
		if b%2 == 1 {
			res = (res * a) % this.mod
		}
		a = (a * a) % this.mod
		b /= 2
	}
	return res
}

func (this *Fancy) Append(val int) {
	// 新元素就是通过反响计算弄回去, 就保证了所有的元素都是可以通过mul, add计算了
	invMul := this.pow(this.mul, this.mod-2)
	newVale := (int64(val) - this.add + this.mod) % this.mod * invMul % this.mod
	this.source = append(this.source, int(newVale))
}

func (this *Fancy) AddAll(inc int) {
	this.add = (this.add + int64(inc)) % this.mod
}

func (this *Fancy) MultAll(m int) {
	this.add = (this.add * int64(m)) % this.mod
	this.mul = (this.mul * int64(m)) % this.mod
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.source) {
		return -1
	}
	v := this.source[idx]
	tmp := v*int(this.mul) + int(this.add)
	return tmp % int(this.mod)
}
