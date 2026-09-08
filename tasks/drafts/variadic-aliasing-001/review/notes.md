# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- If values is empty return false. Otherwise replace only its first element and return true.
- When a caller passes a slice with values..., the caller must observe the change through shared backing storage.
- Preserve length, capacity and all other elements, including storage outside a passed subslice. Zero and negative integers are valid.
- Direct variadic arguments also work. Do not append, clone the slice or mutate a different position.

Public/reference assertions (actual/expected diagnostics):

- Case01: `values:=[]int{1,2,3};if !subject.ReplaceFirst(9,values...)||!slices.Equal(values,[]int{9,2,3}){t.Errorf("expanded values got%v",values)}`
- Case02: `whole:=[]int{7,1,2,8};view:=whole[1:3];beforeCap:=cap(view);if !subject.ReplaceFirst(-4,view...)||!slices.Equal(whole,[]int{7,-4,2,8})||len(view)!=2||cap(view)!=beforeCap{t.Errorf("subslice got%v len%d cap%d",whole,len(view),cap(view))};if subject.ReplaceFirst(1)||subject.ReplaceFirst(1,[]int(nil)...){t.Error("empty accepted")};if !subject.ReplaceFirst(0,2,3){t.Error("direct arguments rejected")}`

Realistic required defect models:

- clone-input
- append-instead
- replace-last
- zero-rejected
- false-after-update

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
