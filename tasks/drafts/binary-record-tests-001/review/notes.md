# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- MarshalBinary returns six bytes: Code as two big-endian bytes followed by Count as four big-endian bytes, with nil error.
- UnmarshalBinary replaces the non-nil receiver from exactly six bytes. Reject every other length with the supplied non-nil ErrRecord and leave the receiver unchanged.
- Preserve the input bytes and sentinel. Returned byte slices from separate marshal calls are independent.
- Support all uint16 and uint32 values; round-trip tests alone cannot establish byte order.

Public/reference assertions (actual/expected diagnostics):

- Case01: `r:=subject.Record{Code:0x1234,Count:0x89abcdef};data,e:=r.MarshalBinary();if e!=nil||!slices.Equal(data,[]byte{0x12,0x34,0x89,0xab,0xcd,0xef}){t.Errorf("wire got%x,%v",data,e)};var got subject.Record;if e=got.UnmarshalBinary([]byte{0x12,0x34,0x89,0xab,0xcd,0xef});e!=nil||got!=r{t.Errorf("decode got%+v,%v",got,e)}`
- Case02: `for _,r:=range []subject.Record{{},{Code:65535,Count:4294967295}}{data,e:=r.MarshalBinary();before:=slices.Clone(data);var got subject.Record;if e!=nil||got.UnmarshalBinary(data)!=nil||got!=r||!slices.Equal(data,before){t.Errorf("boundary got%+v,%x,%v",got,data,e)}}`
- Case03: `sentinel:=subject.ErrRecord;if sentinel==nil{t.Fatal("nil record sentinel")};for _,data:=range [][]byte{nil,{},make([]byte,5),make([]byte,7)}{r:=subject.Record{Code:7,Count:9};before:=slices.Clone(data);e:=r.UnmarshalBinary(data);if e!=sentinel||subject.ErrRecord!=sentinel||r!=(subject.Record{Code:7,Count:9})||!slices.Equal(data,before){t.Errorf("invalid length%d got%+v,%v",len(data),r,e)}}`
- Case04: `r:=subject.Record{Code:1,Count:2};a,e:=r.MarshalBinary();b,f:=r.MarshalBinary();if e!=nil||f!=nil||len(a)!=6||len(b)!=6{t.Fatal("bad marshal results")};a[0]=99;if b[0]!=0{t.Error("marshal outputs share storage")};input:=[]byte{0,1,0,0,0,2};before:=slices.Clone(input);var v subject.Record;if v.UnmarshalBinary(input)!=nil||!slices.Equal(input,before){t.Error("decoder changed input")};input[1]=9;if v.Code!=1{t.Error("receiver retained mutable input")}`

Realistic required defect models:

- little-endian
- drop-high-count
- accept-trailing
- clear-on-error
- mutate-input
- fresh-error
- shared-buffer

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
