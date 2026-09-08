# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Replace the final destination element with the final source element.
- Preserve every earlier destination element and the destination length.
- Changes are visible through existing destination views; preserve every source element.
- Both slices are nonempty and their backing storage does not overlap; lengths may differ.

Public/reference assertions (actual/expected diagnostics):

- Case01: `src:=[]string{"a","b","c","last"};dst:=[]string{"a","x","b","z","?"};subject.CopyLast(src,dst);want:=[]string{"a","x","b","z","last"};if !slices.Equal(dst,want)||!slices.Equal(src,[]string{"a","b","c","last"}){t.Errorf("source/destination got %q/%q want unchanged source/%q",src,dst,want)}`
- Case02: `src:=[]string{""};dst:=[]string{"?"};view:=dst[:];subject.CopyLast(src,dst);if view[0]!=""{t.Errorf("aliased final value got %q want empty",view[0])}`

Realistic required defect models:

- append-instead
- source-first
- destination-first
- copy-destination
- swap-last

Retained source behavior: lasagna-master selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
