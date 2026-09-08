# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the first n values with length and capacity exactly n; preserve their order.
- The result owns independent storage: mutations through input or result must not affect the other.
- Bounds satisfy 0 <= n <= len(in). Nil input and n==0 are valid, with zero capacity.
- Preserve input during the call; no garbage-collector timing or absolute memory measurement is required.

Public/reference assertions (actual/expected diagnostics):

- Case01: `in:=[]int{2,3,4,5};out:=subject.DetachedPrefix(in,2);if !slices.Equal(out,[]int{2,3})||cap(out)!=2||!slices.Equal(in,[]int{2,3,4,5}){t.Fatalf("prefix got%v cap%d input%v",out,cap(out),in)};out[0]=9;if in[0]!=2{t.Error("output shares input")};in[1]=8;if out[1]!=3{t.Error("input shares output")}`
- Case02: `in:=make([]int,2,100);out:=subject.DetachedPrefix(in,0);if len(out)!=0||cap(out)!=0{t.Errorf("empty retained cap%d",cap(out))}`

Realistic required defect models:

- shared-prefix
- capped-but-shared
- copy-entire
- reverse
- empty-retained

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
