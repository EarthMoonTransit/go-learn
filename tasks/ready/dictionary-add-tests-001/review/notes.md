# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- For a missing key store the supplied value and return nil.
- For a present key return the supplied ErrExists sentinel and preserve its previous value.
- A present empty value is still an existing entry; empty keys and new empty values are valid.
- Preserve unrelated entries and ErrExists. The dictionary is non-nil; a fresh error with matching text is not the supplied sentinel.

Public/reference assertions (actual/expected diagnostics):

- Case01: `d:=subject.Dictionary{};e:=d.Add("test","this is just a test");if e!=nil||d["test"]!="this is just a test"{t.Errorf("insert got %v,%v",d,e)}`
- Case02: `d:=subject.Dictionary{"test":"old","other":"keep"};e:=d.Add("test","new");if e!=subject.ErrExists||d["test"]!="old"||d["other"]!="keep"||len(d)!=2{t.Errorf("existing got %v,%v want unchanged,ErrExists",d,e)}`
- Case03: `d:=subject.Dictionary{"empty":""};if e:=d.Add("empty","new");e!=subject.ErrExists||d["empty"]!=""{t.Errorf("present empty got %v,%v",d,e)}`
- Case04: `d:=subject.Dictionary{"other":"keep"};e:=d.Add("","");v,ok:=d[""];if e!=nil||!ok||v!=""||d["other"]!="keep"||len(d)!=2{t.Errorf("empty insertion got %v,%v",d,e)}`

Realistic required defect models:

- overwrite
- fresh-error
- sentinel-nullified
- sentinel-reassigned
- empty-is-absent
- empty-key-rejected
- empty-value-ignored
- clear-others
- success-error

Retained source behavior: lgwt:maps/v7/dictionary.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| new value and nil | 01 | success-error; overwrite |
| existing entry unchanged/sentinel identity | 02 | overwrite; fresh-error |
| original non-nil sentinel preserved | captured identity and post-call assertion in every case | sentinel-nullified; sentinel-reassigned |
| present empty counts as existing | 03 | empty-is-absent |
| empty key/new empty value/unrelated entries | 04 | empty-key-rejected; empty-value-ignored; clear-others |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.

Regression: reference/public tests capture the original non-nil sentinel before each call and assert it remains unchanged afterward. Required nullified/reassigned mutants cannot redefine the expected error by changing a package variable.


## Independent review

The strong learner suite captures the original non-nil ErrExists before calls and checks it remains unchanged; required nullified/reassigned defects close the mutable-expected-error bypass. Existing/new/empty keys and values, sentinel identity and unrelated entries are covered.

Reviewed the actual task contract, starter, public/reference assertions, all 9 required defects and route dependencies. The declared operations are substantive permitted adaptations. Exact source artifact hashes and verbatim MIT notices match the pinned local source checkouts, including Exercism authors/contributors where supplied and the literal LGWT copyright text. The LGWT dictionary adaptation explicitly uses errors.New sentinels while retaining lookup/insertion behavior. Private reference/mutant/review material stays outside learner payloads.

The author-provided targeted report passes, with the strong reference detecting all 9 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full batch gates, allocation reconciliation and T07 acceptance remain separate. Automated fixtures are not learner calibration.
