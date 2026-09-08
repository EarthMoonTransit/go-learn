# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Copy at most limit bytes from src to dst. Return bytes written and the original read/write error; EOF before the limit is successful.
- Return 0,ErrLimit for a negative limit, or 0,nil for zero; neither case may read or write.
- A short write with nil error becomes io.ErrShortWrite. When reading returns bytes and an error together, write the bytes before reporting the error.
- Do not close either stream. Streams are non-nil and follow io contracts; no absolute timing or buffer-size requirement.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{input string;limit int64;want string}{{"abcdef",3,"abc"},{"ab",9,"ab"},{"",2,""}}{var out bytes.Buffer;n,e:=subject.CopyAtMost(&out,strings.NewReader(c.input),c.limit);if e!=nil||n!=int64(len(c.want))||out.String()!=c.want{t.Errorf("copy got%d,%v,%q want%q",n,e,out.String(),c.want)}}`
- Case02: `for _,limit:=range []int64{-1,0}{r:=&trackedReader{data:[]byte("abc")};w:=&trackedWriter{};n,e:=subject.CopyAtMost(w,r,limit);want:=error(nil);if limit<0{want=subject.ErrLimit};if n!=0||e!=want||r.reads!=0||w.writes!=0||r.closes!=0||w.closes!=0{t.Errorf("limit%d got%d,%v reads%d writes%d closes%d,%d",limit,n,e,r.reads,w.writes,r.closes,w.closes)}}`
- Case03: `cause:=errors.New("read failed");r:=&trackedReader{data:[]byte("abc"),err:cause};w:=&trackedWriter{};n,e:=subject.CopyAtMost(w,r,10);if n!=3||e!=cause||string(w.data)!="abc"||r.closes!=0||w.closes!=0{t.Errorf("read error got%d,%v data%q closes%d,%d",n,e,w.data,r.closes,w.closes)}`
- Case04: `for _,cause:=range []error{nil,errors.New("write failed")}{r:=&trackedReader{data:[]byte("abcd")};w:=&trackedWriter{short:true,err:cause};n,e:=subject.CopyAtMost(w,r,10);want:=cause;if want==nil{want=io.ErrShortWrite};if n!=2||e!=want||string(w.data)!="ab"||r.closes!=0||w.closes!=0{t.Errorf("partial write got%d,%v data%q closes%d,%d",n,e,w.data,r.closes,w.closes)}}`

Realistic required defect models:

- copy-all
- copy-exact
- swallow-error
- fresh-limit-error
- close-reader
- close-writer

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
