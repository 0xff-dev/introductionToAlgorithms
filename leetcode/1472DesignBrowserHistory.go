package leetcode

type BrowserHistory struct {
	index, length int
	history       []string
}

func Constructor1472(homepage string) BrowserHistory {
	return BrowserHistory{index: 0, length: 1, history: []string{homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	if this.index == this.length-1 && this.length == len(this.history) {
		this.history = append(this.history, url)
		this.index = this.length
		this.length++
		return
	}
	this.index++
	this.history[this.index] = url
	this.length = this.index + 1
}

func (this *BrowserHistory) Back(steps int) string {
	if this.length == 0 {
		return ""
	}
	this.index -= steps
	if this.index < 0 {
		this.index = 0
	}
	return this.history[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	//a,b,c
	this.index += steps
	if this.index >= this.length {
		this.index = this.length - 1
	}
	return this.history[this.index]
}
