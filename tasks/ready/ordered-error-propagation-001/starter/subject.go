package subject

type Calculator interface {
	Amount(int) (float64, error)
	Factor() (float64, error)
}

func Divide(c Calculator, count int) (float64, error) {
	amount, e := c.Amount(count)
	if e != nil {
	}
	factor, e := c.Factor()
	if e != nil {
		return 0, e
	}
	return amount * factor / float64(count), nil
}
