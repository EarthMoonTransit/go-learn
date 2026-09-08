# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- For an existing key replace its value exactly and return nil; a stored empty value counts as existing.
- For an absent key return ErrAbsent, preserving the entire dictionary; nil maps contain no keys.
- Empty keys and new empty values are valid. Preserve unrelated entries and do not normalize key spelling.
- The failure has dynamic type DictionaryError and Error text "cannot update absent key"; a newly allocated generic error is insufficient.

Public/reference assertions (actual/expected diagnostics):

- Case01: `d:=subject.Dictionary{"test":"old","other":"keep"};e:=d.Update("test","new");if e!=nil||d["test"]!="new"||d["other"]!="keep"||len(d)!=2{t.Errorf("update got %v,%v",d,e)}`
- Case02: `for _,d:=range []subject.Dictionary{nil,{}, {"A":"keep"}}{e:=d.Update("a","new");typed,ok:=e.(subject.DictionaryError);if !ok||typed!=subject.ErrAbsent||e.Error()!="cannot update absent key"{t.Errorf("absent got %T,%v want DictionaryError ErrAbsent",e,e)};if len(d)>0&&(len(d)!=1||d["A"]!="keep"){t.Errorf("absent changed map %v",d)}}`
- Case03: `d:=subject.Dictionary{"":"","a":"v"};if e:=d.Update("","new");e!=nil||d[""]!="new"{t.Errorf("existing empty got %v,%v",d,e)};if e:=d.Update("a","");e!=nil||d["a"]!=""||len(d)!=2{t.Errorf("new empty got %v,%v",d,e)}`

Realistic required defect models:

- insert-missing
- empty-missing
- generic-error
- ignore-value
- clear-others
- empty-key-rejected

Retained source behavior: lgwt:maps/v7/dictionary.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
