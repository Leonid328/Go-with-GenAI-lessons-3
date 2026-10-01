package wallet

type SecureWallet struct {
	balance float64
}

func (w *SecureWallet) Balance() float64 {
	return w.balance
}

func (w *SecureWallet) Deposit(amt float64) {
	w.balance += amt
}

func ApplyDeposits(wallets []SecureWallet, amt float64) {
	for i := range wallets {
		wallets[i].Deposit(amt)
	}
}
