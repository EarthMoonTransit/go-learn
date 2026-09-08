package subject

const rateValue = 5 / 100
const dayValue int = 365
const Rate = rateValue
const Days = dayValue

func Rate32() float32                 { return Rate }
func DayCount() int                   { return Days }
func AnnualFraction(days int) float64 { return float64(days) / float64(Days) }
