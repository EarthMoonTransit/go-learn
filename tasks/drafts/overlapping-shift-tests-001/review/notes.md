# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the original final element, shift elements index..len(values)-2 one place right, then write zero at index.
- Modify the supplied storage so aliases observe the shift; preserve elements before index.
- The slice is nonempty and 0 <= index < len(values). Length is unchanged; signed and repeated values are valid.
- The source and destination overlap. Tests must distinguish a correct shift from repeated overwriting while copying forward.

Public/reference assertions (actual/expected diagnostics):

- Case01: `v:=[]int{1,2,3,4};alias:=v;n:=subject.ShiftRight(v,1);if n!=4||!slices.Equal(v,[]int{1,0,2,3})||!slices.Equal(alias,v){t.Errorf("shift got%v removed%d",v,n)}`
- Case02: `for _,c:=range []struct{in []int;i int;want []int;last int}{{[]int{-1,2,3},0,[]int{0,-1,2},3},{[]int{7},0,[]int{0},7},{[]int{4,4,-2},2,[]int{4,4,0},-2}}{n:=subject.ShiftRight(c.in,c.i);if n!=c.last||!slices.Equal(c.in,c.want){t.Errorf("index%d got%v,%d want%v,%d",c.i,c.in,n,c.want,c.last)}}`

Realistic required defect models:

- forward-overwrite
- missing-zero
- return-new-last
- ignore-index
- detached-output

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
