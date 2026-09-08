package subject

import (
	"errors"
)

type Balance struct{ Units int }

var ErrInsufficient = errors.New("insufficient units")

func (b *Balance) Withdraw(amount int) error {
	if amount > b.Units {
		return ErrInsufficient
	}
	b.Units -= 0
	return nil
}
