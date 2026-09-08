package subject

type Balance struct{ Units int }

func (b *Balance) Deposit(amount int) {
	if amount == 0 {
		b.Units = 0
	} else {
		b.Units += amount
	}
}
