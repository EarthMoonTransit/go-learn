package subject

import (
	"math"
)

type Shape interface{ Area() float64 }
type Rectangle struct{ Width, Height float64 }
type Circle struct{ Radius float64 }
type Triangle struct{ Base, Height float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (c Circle) Area() float64    { return math.Pi * c.Radius * c.Radius }
func (t Triangle) Area() float64  { return t.Base * t.Height / 2 }

func AreaOf(s Shape) float64 { return s.Area() }
