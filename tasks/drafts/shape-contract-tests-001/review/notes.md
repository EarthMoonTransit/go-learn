# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Rectangle area is Width times Height; Circle area is math.Pi times Radius squared; Triangle area is Base times Height divided by two.
- Each value implements Shape without a declaration of implementation. Exercise heterogeneous Shape values in your tests.
- Dimensions are finite from 0 through 1000; allow absolute error 0.000000001 and reject NaN.
- Zero dimensions and fractional dimensions are valid; results do not depend on previous calls.

Public/reference assertions (actual/expected diagnostics):

- Case01: `var shape subject.Shape=subject.Rectangle{Width:12,Height:6};got:=shape.Area();if math.IsNaN(got)||math.Abs(got-72)>1e-9{t.Errorf("rectangle got %g want72",got)}`
- Case02: `cases:=[]struct{s subject.Shape;want float64}{{subject.Circle{Radius:10},math.Pi*100},{subject.Triangle{Base:12,Height:6},36},{subject.Rectangle{Width:.5,Height:3},1.5},{subject.Circle{Radius:.5},math.Pi/4},{subject.Triangle{Base:1,Height:1},.5},{subject.Rectangle{},0},{subject.Circle{},0},{subject.Triangle{Base:4},0}};for i,c:=range cases{got:=c.s.Area();if math.IsNaN(got)||math.Abs(got-c.want)>1e-9{t.Errorf("shape%d got %g want %g",i,got,c.want)}}`

Realistic required defect models:

- rectangle-sum
- circle-linear
- approximate-pi
- triangle-no-half
- fraction-truncated
- zero-fallback

Retained source behavior: lgwt:structs/v8/shapes.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
