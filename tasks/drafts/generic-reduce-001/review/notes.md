# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Starting with initial, call f once per input element in order and carry its returned accumulator to the next call.
- Return initial without calling f on nil or empty input. The callback is non-nil.
- Input and result types may differ. Preserve the input slice; no concurrent calls or environment access.
- Support named types and non-comparable accumulators; arithmetic inputs in tests do not overflow.

Public/reference assertions (actual/expected diagnostics):

- Case01: `in:=[]int{1,2,3};calls:=[]int{};got:=subject.Reduce(in,func(s string,n int)string{calls=append(calls,n);return fmt.Sprintf("%s%d",s,n)},"x");if got!="x123"||!slices.Equal(calls,[]int{1,2,3})||!slices.Equal(in,[]int{1,2,3}){t.Errorf("ordered reduce got%q calls%v input%v",got,calls,in)}`
- Case02: `calls:=0;got:=subject.Reduce[int,string](nil,func(s string,n int)string{calls++;return "bad"},"initial");if got!="initial"||calls!=0{t.Errorf("empty got%q calls%d",got,calls)};v:=subject.Reduce([]int{2,3},func(a,b int)int{return a-b},10);if v!=5{t.Errorf("carry initial got%d want5",v)}`
- Case03: `type count int;got:=subject.Reduce([]count{1,2},func(a []int,v count)[]int{return append(a,int(v))},[]int{9});if !slices.Equal(got,[]int{9,1,2}){t.Errorf("noncomparable accumulator got%v",got)}`

Realistic required defect models:

- reset-each-step
- zero-initial
- skip-first
- reverse-order
- double-step
- erase-input

Retained source behavior: lgwt:arrays/v8/collection_fun.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
