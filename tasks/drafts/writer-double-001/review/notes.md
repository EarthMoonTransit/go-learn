# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- NewWriter creates an empty independent writer with a nonnegative total byte limit and a supplied non-nil failure error.
- Each Write records one call and accepts the prefix fitting the remaining total capacity. If all bytes fit return count,nil, otherwise count and the original failure.
- An empty Write still counts and succeeds. Reaching the limit exactly succeeds; later nonempty writes fail with zero accepted bytes.
- Snapshot returns recorded bytes as an independent copy and the call count. Do not retain mutable input storage or share state between instances.

Public/reference assertions (actual/expected diagnostics):

- Case01: `cause:=errors.New("limit");w:=subject.NewWriter(5,cause);if w==nil{t.Fatal("nil writer")};n,e:=w.Write([]byte("ab"));if n!=2||e!=nil{t.Errorf("initial got%d,%v",n,e)};n,e=w.Write([]byte("cdef"));data,calls:=w.Snapshot();if n!=3||e!=cause||string(data)!="abcde"||calls!=2{t.Errorf("partial got%d,%v data%q calls%d",n,e,data,calls)};n,e=w.Write([]byte("x"));if n!=0||e!=cause{t.Errorf("full got%d,%v",n,e)};n,e=w.Write(nil);_,calls=w.Snapshot();if n!=0||e!=nil||calls!=4{t.Errorf("empty got%d,%v calls%d",n,e,calls)}`
- Case02: `cause:=errors.New("limit");a:=subject.NewWriter(2,cause);b:=subject.NewWriter(0,cause);if a==nil||b==nil{t.Fatal("nil writer")};input:=[]byte("ab");n,e:=a.Write(input);if n!=2||e!=nil{t.Errorf("exact limit got%d,%v",n,e)};input[0]='x';data,calls:=a.Snapshot();if string(data)!="ab"||calls!=1{t.Errorf("input alias data%q calls%d",data,calls)};data[0]='z';next,_:=a.Snapshot();if string(next)!="ab"{t.Errorf("snapshot alias %q",next)};if n,e:=b.Write([]byte("x"));n!=0||e!=cause{t.Errorf("zero limit got%d,%v",n,e)};next,calls=b.Snapshot();if len(next)!=0||calls!=1{t.Errorf("independent writer %q calls%d",next,calls)}`

Realistic required defect models:

- per-call-limit
- error-at-exact-limit
- fresh-error
- no-call-count
- snapshot-alias
- retain-input

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
