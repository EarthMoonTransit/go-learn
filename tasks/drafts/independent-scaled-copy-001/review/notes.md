# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Input quantities represent two units; multiply each by units divided by two.
- Preserve order and length, including zero and fractional quantities.
- Return independent output storage: changing it later must not change input values.
- Input values are finite between 0 and 10000, units is 0 through 100; preserve input even for units 0 or 2.
- Nil or empty input returns zero values; allow absolute floating error 0.000000001.

Public/reference assertions (actual/expected diagnostics):

- Case01: `in:=[]float64{0.5,250,150,3,0.5};want:=[]float64{1.5,750,450,9,1.5};got:=subject.ScaledCopy(in,6);if !slices.Equal(got,want){t.Errorf("scaled got %v want %v",got,want)};if !slices.Equal(in,[]float64{0.5,250,150,3,0.5}){t.Errorf("input changed: %v",in)}`
- Case02: `got:=subject.ScaledCopy([]float64{0.6,1,0.5},3);want:=[]float64{0.9,1.5,0.75};if len(got)!=len(want){t.Fatalf("length got %d want3",len(got))};for i:=range want{if math.IsNaN(got[i])||math.Abs(got[i]-want[i])>1e-9{t.Errorf("value%d got %g want %g",i,got[i],want[i])}}`
- Case03: `for _,units:=range []int{0,2}{in:=[]float64{4,2};got:=subject.ScaledCopy(in,units);if len(got)!=2{t.Fatalf("length got %d want2",len(got))};got[0]=99;if in[0]!=4||in[1]!=2{t.Errorf("units%d shared or modified input %v",units,in)}}`
- Case04: `if got:=subject.ScaledCopy(nil,100);len(got)!=0{t.Errorf("nil got %v want empty",got)}`

Realistic required defect models:

- in-place
- integer-factor
- missing-base
- same-scale-alias
- zero-scale-alias
- reverse-output
- empty-placeholder

Retained source behavior: lasagna-master selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| scaling/length/order | 01–02 | missing-base; integer-factor; reverse-output |
| input preserved | 01,03 | in-place; zero-scale-alias |
| independent zero/unchanged-scale storage | 03 | same-scale-alias; zero-scale-alias |
| zero scale produces zero values | 03 | zero-scale-ignored |
| nil input | 04 | empty-placeholder |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.
