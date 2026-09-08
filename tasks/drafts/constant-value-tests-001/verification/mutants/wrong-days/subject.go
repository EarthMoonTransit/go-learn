package subject

const rateValue = 0.05
const dayValue int = 360
const Rate = rateValue
const Days = dayValue

func Rate32() float32                 { return Rate }
func DayCount() int                   { return Days }
func AnnualFraction(days int) float64 { return float64(days) / float64(Days) }
