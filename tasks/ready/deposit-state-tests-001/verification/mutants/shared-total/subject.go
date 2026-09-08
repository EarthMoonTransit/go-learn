package subject

type Balance struct{ Units int }

var total int

func (b *Balance) Deposit(amount int) { total += amount; b.Units = total }
