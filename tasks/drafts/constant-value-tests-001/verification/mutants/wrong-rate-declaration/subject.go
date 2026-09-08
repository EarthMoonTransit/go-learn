package subject

const rateValue = 0.1
const dayValue int = 365
const Rate = rateValue
const Days = dayValue

func Rate32() float32                 { return .05 }
func DayCount() int                   { return 365 }
func AnnualFraction(days int) float64 { return float64(days) / 365 }
