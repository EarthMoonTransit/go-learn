package subject

const rateValue = 0.05
const dayValue int = 360
const Rate = rateValue
const Days = dayValue

func Rate32() float32                 { return .05 }
func DayCount() int                   { return 365 }
func AnnualFraction(days int) float64 { return float64(days) / 365 }
