package leetcode

/*
	type Bitset struct {
		bits []uint8
		size int
		one  int

		binTable [256]string
	}

	func Constructor2166(size int) Bitset {
		n := size / 8
		if size%8 != 0 {
			n++
		}
		b := Bitset{
			bits:     make([]uint8, n),
			size:     size,
			one:      0,
			binTable: [256]string{},
		}
		for i := range 256 {
			b.binTable[i] = fmt.Sprintf("%08b", uint8(i))
		}
		return b
	}

	func (this *Bitset) Fix(idx int) {
		index := idx / 8
		step := idx % 8
		mask := uint8(1) << (7 - step)
		// 0 0 0 0 0 0 0 0
		//       1
		if this.bits[index]&mask == mask {
			return
		}
		this.bits[index] |= mask
		this.one++
	}

	func (this *Bitset) Unfix(idx int) {
		index := idx / 8
		step := idx % 8
		mask := ^(uint8(1) << (7 - step))
		// 0 0 0 0 0 0 0 0
		//       1
		if this.bits[index]&mask == this.bits[index] {
			return
		}
		this.bits[index] &= mask
		this.one--
	}

	func (this *Bitset) Flip() {
		for i := range this.bits {
			this.bits[i] = ^this.bits[i]
		}
		this.one = this.size - this.one
	}

	func (this *Bitset) All() bool {
		return this.one == this.size
	}

	func (this *Bitset) One() bool {
		return this.one > 0
	}

	func (this *Bitset) Count() int {
		return this.one
	}

	func (this *Bitset) ToString() string {
		var buf strings.Builder
		for i := 0; i < len(this.bits)-1; i++ {
			buf.WriteString(this.binTable[this.bits[i]])
		}
		last := this.binTable[this.bits[len(this.bits)-1]]
		if step := this.size % 8; step != 0 {
			last = last[:step]
		}
		buf.WriteString(last)
		return buf.String()
	}
*/
type Bitset struct {
	bits      []uint64
	size, one int

	flipped uint64
}

func Constructor2166(size int) Bitset {
	return Bitset{
		bits: make([]uint64, (size+63)/64),
		size: size,
	}
}

func (this *Bitset) Fix(idx int) {
	i, m := idx>>6, uint64(1)<<(idx&63)
	if (this.bits[i]^this.flipped)&m == 0 {
		this.bits[i] ^= m
		this.one++
	}
}

func (this *Bitset) Unfix(idx int) {
	i, m := idx>>6, uint64(1)<<(idx&63)
	if (this.bits[i]^this.flipped)&m != 0 {
		this.bits[i] ^= m
		this.one--
	}
}

func (this *Bitset) Flip() {
	this.flipped = ^this.flipped
	this.one = this.size - this.one
}

func (this *Bitset) All() bool  { return this.one == this.size }
func (this *Bitset) One() bool  { return this.one > 0 }
func (this *Bitset) Count() int { return this.one }

func (this *Bitset) ToString() string {
	buf := make([]byte, this.size)
	for i := range buf {
		w := this.bits[i>>6] ^ this.flipped
		buf[i] = '0' + byte(w>>(i&63)&1)
	}
	return string(buf)
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
