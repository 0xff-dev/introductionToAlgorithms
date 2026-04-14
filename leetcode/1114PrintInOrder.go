package leetcode

type Foo struct {
	b, c chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		b: make(chan struct{}, 1),
		c: make(chan struct{}, 1),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.b <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
	<-f.b
	printSecond()
	f.c <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	<-f.c
	printThird()
}
