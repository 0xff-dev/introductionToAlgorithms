package leetcode

type H2O struct {
	h, o chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
		h: make(chan struct{}, 2),
		o: make(chan struct{}, 1),
	}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.h <- struct{}{}
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.o <- struct{}{}
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	<-h.h
	<-h.h
	<-h.o
}
