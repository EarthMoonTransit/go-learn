# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Read root files in lexical name order, skipping directories. Each starts with Title: , Description: , Tags: headers and a fourth line ---.
- Strip those exact prefixes, split nonempty tags on comma-space, and join remaining ScanLines body lines with newline. Preserve other spaces; empty tags/body are allowed. Lines are at most 2048 bytes.
- Close every successfully opened file exactly once before returning. Read/format errors take precedence over close errors; otherwise return the original close error.
- Any directory/open/read/format/close failure returns nil records and the original error; malformed headers use the supplied non-nil ErrFormat. Preserve that sentinel. Empty directories return no records.

Public/reference assertions (actual/expected diagnostics):

- Case01: `base:=fstest.MapFS{"a.md":{Data:[]byte("Title: Post 1\nDescription: Description 1\nTags: tdd, go\n---\nHello\nWorld")},"b.md":{Data:[]byte("Title: B\nDescription: D\nTags: rust, borrow-checker\n---\nB\nL\nM")},"sub/file":{Data:[]byte("ignored")}};fsys:=&trackedFS{base:base};got,e:=subject.ReadRecords(fsys);want:=[]subject.Post{{Title:"Post 1",Description:"Description 1",Tags:[]string{"tdd","go"},Body:"Hello\nWorld"},{Title:"B",Description:"D",Tags:[]string{"rust","borrow-checker"},Body:"B\nL\nM"}};if e!=nil||!reflect.DeepEqual(got,want)||!slices.Equal(fsys.opened,[]string{"a.md","b.md"})||!slices.Equal(fsys.closed,fsys.opened){t.Errorf("records got%+v,%v opens%v closes%v",got,e,fsys.opened,fsys.closed)}`
- Case02: `fsys:=&trackedFS{base:fstest.MapFS{"x":{Data:[]byte("Title:  x \nDescription: \nTags: \n---\n a \n\nλ\n")}}};got,e:=subject.ReadRecords(fsys);if e!=nil||len(got)!=1||got[0].Title!=" x "||got[0].Description!=""||len(got[0].Tags)!=0||got[0].Body!=" a \n\nλ"||len(fsys.closed)!=1{t.Errorf("exact fields got%+v,%v closes%v",got,e,fsys.closed)};got,e=subject.ReadRecords(&trackedFS{base:fstest.MapFS{}});if e!=nil||len(got)!=0{t.Errorf("empty directory got%v,%v",got,e)}`
- Case03: `cause:=errors.New("filesystem failed");valid:=[]byte("Title: T\nDescription: D\nTags: go\n---\nbody");for _,kind:=range []string{"directory","open","read","close"}{fsys:=&trackedFS{base:fstest.MapFS{"a":{Data:valid},"b":{Data:valid}},failName:"b"};switch kind{case "directory":fsys.dirErr=cause;case "open":fsys.openErr=cause;case "read":fsys.readErr=cause;case "close":fsys.closeErr=cause};got,e:=subject.ReadRecords(fsys);if got!=nil||e!=cause||!slices.Equal(fsys.closed,fsys.opened){t.Errorf("%s got%v,%v opens%v closes%v",kind,got,e,fsys.opened,fsys.closed)}}`
- Case04: `sentinel:=subject.ErrFormat;if sentinel==nil{t.Fatal("nil format sentinel")};for _,body:=range []string{"Title: T\nDescription: D\nTags: go\nwrong","Title:T\nDescription: D\nTags: go\n---","Title: T\nDescription:D\nTags: go\n---","Title: T\nDescription: D\nTags:go\n---","Title: T"}{fsys:=&trackedFS{base:fstest.MapFS{"x":{Data:[]byte(body)}},failName:"x",closeErr:errors.New("close failed")};got,e:=subject.ReadRecords(fsys);if got!=nil||e!=sentinel||subject.ErrFormat!=sentinel||!slices.Equal(fsys.closed,[]string{"x"}){t.Errorf("malformed got%v,%v closes%v",got,e,fsys.closed)}};readErr:=errors.New("read first");fsys:=&trackedFS{base:fstest.MapFS{"x":{Data:[]byte("partial")}},failName:"x",readErr:readErr,closeErr:errors.New("close second")};got,e:=subject.ReadRecords(fsys);if got!=nil||e!=readErr||!slices.Equal(fsys.closed,[]string{"x"}){t.Errorf("error priority got%v,%v closes%v",got,e,fsys.closed)}`

Realistic required defect models:

- no-close
- ignore-close-error
- partial-results
- reverse-files
- ignore-reader-error
- unsplit-tags
- trim-body
- fresh-open-error
- close-before-read

Retained source behavior: lgwt:reading-files/blogposts.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
