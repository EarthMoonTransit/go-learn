package subject

type Balance struct{ Units int }

func (b *Balance) Deposit(amount int) { b.Units += amount }
