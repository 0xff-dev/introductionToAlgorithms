package leetcode

type ATM struct {
	value [5]int
	atm   [5]int
}

func Constructor2241() ATM {
	return ATM{
		value: [5]int{20, 50, 100, 200, 500},
		atm:   [5]int{},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := range 5 {
		this.atm[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	index := 4
	ret := make([]int, 5)
	for ; index >= 0; index-- {
		if this.atm[index] != 0 {
			break
		}
	}
	if index == -1 {
		// 比如19的情况
		return []int{-1}
	}
	source := this.atm
	for ; index >= 0 && amount > 0; index-- {
		if this.atm[index] > 0 {
			c := amount / this.value[index]
			use := min(c, this.atm[index])
			amount -= use * this.value[index]
			this.atm[index] = max(0, this.atm[index]-use)
			ret[index] = use
		}
	}
	if amount > 0 {
		// 不要见
		this.atm = source
		return []int{-1}
	}
	return ret
}
