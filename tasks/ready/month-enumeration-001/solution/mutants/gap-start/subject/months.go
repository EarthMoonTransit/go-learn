package subject

const (
	jan = iota + 2
	feb
	mar
	apr
	may
	jun
	jul
	aug
	sep
	oct
	nov
	dec
)

func GetMonth(month int) int { return month }
