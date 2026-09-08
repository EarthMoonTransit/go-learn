# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the first matching value and true, or the zero value of T and false when no value matches.
- Call the non-nil predicate once per visited element in input order, stopping at the first match; never call it for empty input.
- Preserve the input slice. The function works with any element type, including non-comparable structs containing slices.
- Use inferred and explicit type arguments in tests; a matched zero value still has found=true.

Public/reference assertions (actual/expected diagnostics):

- Case01: `v,ok:=subject.Find([]int{1,2,4},func(x int)bool{return x%2==0});if v!=2||!ok{t.Errorf("first even got%d,%v",v,ok)}`
- Case02: `visited:=[]int{};in:=[]int{3,0,5};v,ok:=subject.Find(in,func(x int)bool{visited=append(visited,x);return x==0});if v!=0||!ok||!slices.Equal(visited,[]int{3,0})||!slices.Equal(in,[]int{3,0,5}){t.Errorf("zero/visits/preservation got%d,%v %v %v",v,ok,visited,in)}`
- Case03: `calls:=0;v,ok:=subject.Find[string](nil,func(x string)bool{calls++;return true});if v!=""||ok||calls!=0{t.Errorf("empty got%q,%v calls%d",v,ok,calls)};v,ok=subject.Find([]string{"a","b"},func(string)bool{return false});if v!=""||ok{t.Errorf("no match got%q,%v",v,ok)}`
- Case04: `type row struct{Data []int};v,ok:=subject.Find([]row{{[]int{1}},{[]int{2,3}}},func(x row)bool{return len(x.Data)==2});if !ok||!slices.Equal(v.Data,[]int{2,3}){t.Errorf("noncomparable got%v,%v",v,ok)}`

Realistic required defect models:

- last-match
- first-unchecked
- false-success
- zero-on-success
- true-miss
- erase-input
- double-predicate

Retained source behavior: lgwt:arrays/v8/collection_fun.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
