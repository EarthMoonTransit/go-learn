package subject

type Balance struct{ Units int }

var previous *Balance

func (b *Balance) Deposit(amount int) {
	if previous != nil && previous != b {
		previous.Units += amount
	}
	b.Units += amount
	previous = b
}
