# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the exact stored value and nil for a present key, including a present empty value.
- For an absent key return empty text and the supplied ErrMissing sentinel.
- Nil and empty dictionaries have no keys; keys are exact strings and may be empty.
- Preserve all dictionary entries. A newly created error with the same text does not equal the supplied sentinel.

Public/reference assertions (actual/expected diagnostics):

- Case01: `d:=subject.Dictionary{"test":"this is just a test"};v,e:=d.Search("test");if v!="this is just a test"||e!=nil{t.Errorf("present got (%q,%v)",v,e)}`
- Case02: `v,e:=subject.Dictionary{"key":""}.Search("key");if v!=""||e!=nil{t.Errorf("present empty got (%q,%v) want empty,nil",v,e)}`
- Case03: `v,e:=subject.Dictionary{"test":"value"}.Search("unknown");if v!=""||e!=subject.ErrMissing{t.Errorf("absent got (%q,%v) want empty,ErrMissing",v,e)}`
- Case04: `var d subject.Dictionary;v,e:=d.Search("");if v!=""||e!=subject.ErrMissing{t.Errorf("nil map got (%q,%v) want empty,ErrMissing",v,e)}`
- Case05: `d:=subject.Dictionary{"":"v","A":"x"};v,e:=d.Search("");if v!="v"||e!=nil{t.Errorf("empty key got (%q,%v)",v,e)};_,e=d.Search("a");if e!=subject.ErrMissing{t.Errorf("case lookup got %v want ErrMissing",e)};if len(d)!=2||d[""]!="v"||d["A"]!="x"{t.Errorf("dictionary changed: %v",d)}`

Realistic required defect models:

- empty-is-missing
- fresh-error
- missing-nil
- value-erased
- empty-key-rejected
- destructive-search
- case-normalized

Retained source behavior: lgwt:maps/v7/dictionary.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| present exact value and nil | 01–02 | value-erased; empty-is-missing |
| absent stable sentinel | 03–04 | fresh-error; missing-nil |
| exact empty/case keys and preservation | 05 | empty-key-rejected; case-normalized; destructive-search |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.

Regression: reference/public tests capture the original non-nil sentinel before each call and assert it remains unchanged afterward. Required nullified/reassigned mutants cannot redefine the expected error by changing a package variable.
