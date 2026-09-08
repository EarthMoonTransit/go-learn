# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return ten integers 0..9 in order, five strings "a".."e", and a view containing strings at indices one and two.
- The view shares storage with the five-string slice; view index zero corresponds to original index one.
- Separate calls return independent number and string storage.
- Check values and lengths; do not require an implementation-dependent capacity for the appended integer slice.

Public/reference assertions (actual/expected diagnostics):

- Case01: `n,_,_:=subject.Sequences();if !slices.Equal(n,[]int{0,1,2,3,4,5,6,7,8,9}){t.Errorf("numbers got%v",n)}`
- Case02: `n,w,v:=subject.Sequences();m,x,y:=subject.Sequences();if !slices.Equal(w,[]string{"a","b","c","d","e"})||!slices.Equal(v,[]string{"b","c"})||!slices.Equal(x,w)||!slices.Equal(y,v)||len(n)!=10||len(m)!=10{t.Fatalf("sequences n%v w%v v%v",n,w,v)};v[0]="changed";if w[1]!="changed"||w[0]!="a"||x[1]!="b"||y[0]!="b"{t.Error("subslice alias or independence incorrect")};w[2]="back";if v[1]!="back"{t.Error("reverse alias missing")};n[0]=99;if m[0]!=0{t.Error("number storage shared")}`

Realistic required defect models:

- missing-last
- wrong-start
- short-view
- detached-view
- shared-numbers
- shared-words

Retained source behavior: ardan:topics/go/language/slices/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
