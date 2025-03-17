package leetcode

type Bank struct {
	b  []int64
	ok int
}

func Constructor2043(balance []int64) Bank {
	b := make([]int64, len(balance)+1)
	for i := 1; i <= len(balance); i++ {
		b[i] = balance[i-1]
	}
	return Bank{b: b, ok: len(balance)}
}

func (this *Bank) validateAccount(a int) bool {
	return a >= 1 && a <= this.ok
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.validateAccount(account1) || !this.validateAccount(account2) {
		return false
	}
	have := this.b[account1]
	if have >= money {
		this.b[account1] -= money
		this.b[account2] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.validateAccount(account) {
		return false
	}
	this.b[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	// 0 ,1, 2, 3
	if !this.validateAccount(account) {
		return false
	}
	have := this.b[account]
	if have >= money {
		this.b[account] -= money
		return true
	}
	return false
}
