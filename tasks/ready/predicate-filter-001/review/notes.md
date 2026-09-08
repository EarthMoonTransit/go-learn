# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Call the supplied non-nil predicate once per input record in input order.
- Return every record whose predicate result is true, preserving order and all fields.
- Do not change input records; repeated equal records remain repeated.
- Nil and empty inputs return zero records without invoking the predicate; nil or empty output is acceptable.

Public/reference assertions (actual/expected diagnostics):

- Case01: `in:=[]subject.Record{{Day:1,Amount:5.15,Category:"a"},{Day:13,Amount:55.67,Category:"b"},{Day:15,Amount:11,Category:"a"}};got:=subject.Filter(in,func(r subject.Record)bool{return r.Category=="a"});want:=[]subject.Record{in[0],in[2]};if !slices.Equal(got,want){t.Errorf("filtered got %+v want %+v",got,want)}`
- Case02: `in:=[]subject.Record{{Day:3},{Day:1},{Day:3}};var called []int;got:=subject.Filter(in,func(r subject.Record)bool{called=append(called,r.Day);return true});if !slices.Equal(called,[]int{3,1,3})||!slices.Equal(got,in){t.Errorf("order/calls got %+v,%v want %+v,[3 1 3]",got,called,in)}`
- Case03: `in:=[]subject.Record{{Day:1},{Day:2}};before:=append([]subject.Record(nil),in...);subject.Filter(in,func(r subject.Record)bool{return r.Day==2});if !slices.Equal(in,before){t.Errorf("input changed: %+v",in)}`
- Case04: `calls:=0;got:=subject.Filter(nil,func(r subject.Record)bool{calls++;return true});if len(got)!=0||calls!=0{t.Errorf("empty got %+v calls%d want empty/0",got,calls)}`
- Case05: `got:=subject.Filter([]subject.Record{{Day:1}},func(r subject.Record)bool{return false});if len(got)!=0{t.Errorf("reject all got %+v",got)}`

Realistic required defect models:

- predicate-inverted
- first-only
- reverse-order
- double-call
- compact-input
- drop-fields

Retained source behavior: expenses selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| selected records/all fields | 01 | predicate-inverted; drop-fields |
| order/repeats/exactly one predicate call | 02 | first-only; reverse-order; double-call |
| input unchanged | 03 | compact-input |
| empty/no matches | 04–05 | predicate-inverted; first-only |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Public checks observe predicate call count/order, selected output order and fields, repeated records, input preservation and empty/no-match behavior.

Reviewed the prompt, API, actual checks, all 6 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
