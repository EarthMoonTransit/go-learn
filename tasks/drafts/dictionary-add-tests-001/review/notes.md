# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- For a missing key store the supplied value and return nil.
- For a present key return the supplied ErrExists sentinel and preserve its previous value.
- A present empty value is still an existing entry; empty keys and new empty values are valid.
- Preserve unrelated entries. The dictionary is non-nil; a fresh error with matching text is not the supplied sentinel.

Public/reference assertions (actual/expected diagnostics):

- Case01: `d:=subject.Dictionary{};e:=d.Add("test","this is just a test");if e!=nil||d["test"]!="this is just a test"{t.Errorf("insert got %v,%v",d,e)}`
- Case02: `d:=subject.Dictionary{"test":"old","other":"keep"};e:=d.Add("test","new");if e!=subject.ErrExists||d["test"]!="old"||d["other"]!="keep"||len(d)!=2{t.Errorf("existing got %v,%v want unchanged,ErrExists",d,e)}`
- Case03: `d:=subject.Dictionary{"empty":""};if e:=d.Add("empty","new");e!=subject.ErrExists||d["empty"]!=""{t.Errorf("present empty got %v,%v",d,e)}`
- Case04: `d:=subject.Dictionary{"other":"keep"};e:=d.Add("","");v,ok:=d[""];if e!=nil||!ok||v!=""||d["other"]!="keep"||len(d)!=2{t.Errorf("empty insertion got %v,%v",d,e)}`

Realistic required defect models:

- overwrite
- fresh-error
- empty-is-absent
- empty-key-rejected
- empty-value-ignored
- clear-others
- success-error

Retained source behavior: lgwt:maps/v7/dictionary.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
