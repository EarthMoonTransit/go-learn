# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return an ordinary value copy of input: Label and Count are copied and Items retains the same slice view.
- Changing copied scalar fields does not affect the caller. Changing an existing Items element is visible through both structs.
- Preserve nil versus allocated-empty slices and the original slice length and capacity.
- Do not deep-copy, normalize or alter the input. Do not compare a struct containing a slice with ==.

Public/reference assertions (actual/expected diagnostics):

- Case01: `input:=subject.Snapshot{Label:" x ",Count:7,Items:[]int{1,2}};got:=subject.CopySnapshot(input);if got.Label!=input.Label||got.Count!=7||!slices.Equal(got.Items,input.Items){t.Fatalf("copy got%+v",got)};got.Label="other";got.Count=9;if input.Label!=" x "||input.Count!=7{t.Error("scalar values changed")};got.Items[0]=8;if input.Items[0]!=8{t.Error("slice backing storage not shared")};input.Items[1]=6;if got.Items[1]!=6{t.Error("reverse sharing missing")}`
- Case02: `for _,items:=range [][]int{nil,make([]int,0),make([]int,0,4),make([]int,2,5)}{got:=subject.CopySnapshot(subject.Snapshot{Items:items});if (got.Items==nil)!=(items==nil)||len(got.Items)!=len(items)||cap(got.Items)!=cap(items){t.Errorf("slice header got nil%v len%d cap%d want nil%v len%d cap%d",got.Items==nil,len(got.Items),cap(got.Items),items==nil,len(items),cap(items))}}`

Realistic required defect models:

- deep-copy
- clear-count
- trim-label
- drop-capacity
- nil-to-empty
- destructive-items

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
