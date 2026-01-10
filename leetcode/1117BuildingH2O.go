package leetcode

type H2O struct {
	hChan chan struct{}
	oChan chan struct{}
	done  chan struct{}
}

func NewH2O() *H2O {
	return &H2O{
		hChan: make(chan struct{}, 2), // 允许 2 个 H 进入
		oChan: make(chan struct{}, 1), // 允许 1 个 O 进入
		done:  make(chan struct{}),    // 用于三者同步
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.hChan <- struct{}{}

	releaseHydrogen()
	h.done <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.oChan <- struct{}{}

	releaseOxygen()

	<-h.done
	<-h.done

	<-h.hChan
	<-h.hChan
	<-h.oChan
}
