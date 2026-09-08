package subject

type Balance struct{ Units int }

func (b *Balance) Deposit(amount int) { _ = amount }
