# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return an independent copy with length len(in) and capacity len(in)+extra.
- Preserve input values and order; later output changes and appends must not affect input.
- Extra is 0..1000, input length <=1000. Nil and empty input are valid, including positive extra capacity.
- The exact capacity is observable API behavior, not a timing or allocation benchmark.

Public/reference assertions (actual/expected diagnostics):

- Case01: `in:=[]int{3,-2,0};out:=subject.ReservedCopy(in,4);if !slices.Equal(out,in)||len(out)!=3||cap(out)!=7{t.Fatalf("copy got%v len%d cap%d",out,len(out),cap(out))};out[0]=9;out=append(out,8);if !slices.Equal(in,[]int{3,-2,0}){t.Errorf("input changed %v",in)}`
- Case02: `for _,extra:=range []int{0,3}{out:=subject.ReservedCopy(nil,extra);if len(out)!=0||cap(out)!=extra{t.Errorf("empty extra%d len%d cap%d",extra,len(out),cap(out))}}`

Realistic required defect models:

- extra-as-length
- no-reserve
- missing-copy
- zero-extra-alias
- empty-no-reserve

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
