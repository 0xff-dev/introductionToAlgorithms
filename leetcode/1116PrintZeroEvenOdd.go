package leetcode

type ZeroEvenOdd struct {
	n               int
	index           int
	zero, even, odd chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:     n,
		index: 1,
		zero:  make(chan struct{}, 1),
		even:  make(chan struct{}, 1),
		odd:   make(chan struct{}, 1),
	}
	zeo.zero <- struct{}{}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for range z.zero {
		printNumber(0)
		if z.index&1 == 1 {
			z.odd <- struct{}{}
		} else {
			z.even <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for range z.even {
		printNumber(z.index)
		z.index++
		if z.index <= z.n {
			z.zero <- struct{}{}
			continue
		}
		close(z.zero)
		close(z.even)
		close(z.odd)
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for range z.odd {
		printNumber(z.index)
		z.index++
		if z.index <= z.n {
			z.zero <- struct{}{}
			continue
		}
		close(z.zero)
		close(z.even)
		close(z.odd)
	}
}
