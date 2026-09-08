package subject

type Rectangle struct{ Width, Height float64 }

func Perimeter(rect Rectangle) float64 {
	return 2 * (float64(int(rect.Width)) + float64(int(rect.Height)))
}
