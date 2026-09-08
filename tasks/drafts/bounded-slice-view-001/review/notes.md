# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the elements in in[start:end] with length and capacity both end-start.
- Existing view elements share input storage, so changing one changes the corresponding input element.
- Appending an element to the returned view must allocate independent storage and preserve input elements outside the view.
- Bounds satisfy 0 <= start <= end <= len(in). Empty ranges and nil input are valid; do not mutate input during the call.

Public/reference assertions (actual/expected diagnostics):

- Case01: `in:=[]int{1,2,3,4};v:=subject.CappedView(in,1,3);if !slices.Equal(v,[]int{2,3})||len(v)!=2||cap(v)!=2{t.Fatalf("view got%v len%d cap%d",v,len(v),cap(v))};v[0]=9;if in[1]!=9{t.Errorf("view detached: %v",in)};v=append(v,8);if in[3]!=4{t.Errorf("append overwrote outside: %v",in)}`
- Case02: `for _,bounds:=range [][2]int{{0,0},{2,2},{3,3}}{in:=[]int{1,2,3};v:=subject.CappedView(in,bounds[0],bounds[1]);if len(v)!=0||cap(v)!=0||!slices.Equal(in,[]int{1,2,3}){t.Errorf("empty %v len%d cap%d input%v",bounds,len(v),cap(v),in)}}`

Realistic required defect models:

- two-index
- detached
- ignore-start
- empty-tail-cap
- mutate-input

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
