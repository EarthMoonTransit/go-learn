package subject

import (
	"errors"
)

type Balance struct{ Units int }

var ErrInsufficient = errors.New("insufficient units")

func (b *Balance) Withdraw(amount int) error { return nil }
