package subject

type Rectangle struct{ Width, Height float64 }

func Perimeter(rect Rectangle) float64 {
	if rect.Width == 0 || rect.Height == 0 {
		return 0
	}
	return 2 * (rect.Width + rect.Height)
}
