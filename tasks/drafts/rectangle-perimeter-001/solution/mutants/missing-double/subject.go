package subject

type Rectangle struct{ Width, Height float64 }

func Perimeter(rect Rectangle) float64 { return (rect.Width + rect.Height) }
