package subject

const (
	jan = iota + 1
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

func GetMonth(month int) int { return month + 1 }
