# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Write tests in package subject_test using only NewWindow and Contains; the private fields are not part of the external API.
- Contains reports membership in [min,max); constructor inputs satisfy min<=max and may be negative.
- An empty window contains nothing. Separate windows retain independent bounds.
- Test the public behavior at both boundaries without reproducing or reading its private storage.

Public/reference assertions (actual/expected diagnostics):

- Case01: `w:=subject.NewWindow(2,5);if !w.Contains(3){t.Error("interior missing")}`
- Case02: `for _,c:=range []struct{min,max int}{{2,5},{-3,0},{0,0},{7,7}}{w:=subject.NewWindow(c.min,c.max);for _,v:=range []int{c.min-1,c.min,c.max-1,c.max,c.max+1}{want:=v>=c.min&&v<c.max;if got:=w.Contains(v);got!=want{t.Errorf("window[%d,%d) value%d got%v want%v",c.min,c.max,v,got,want)}}}`
- Case03: `a:=subject.NewWindow(1,3);b:=subject.NewWindow(7,9);if !a.Contains(2)||a.Contains(8)||!b.Contains(8)||b.Contains(2){t.Error("independent windows changed")};empty:=subject.NewWindow(4,4);if empty.Contains(4)||!a.Contains(2){t.Error("empty constructor changed existing window")}`

Realistic required defect models:

- include-upper
- exclude-lower
- wrong-lower
- empty-is-point
- swapped-bounds

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
