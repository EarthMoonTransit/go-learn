# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Return a non-nil function that accepts a record when Day is between From and To inclusive.
- For From greater than To, every record is rejected.
- Only Day affects acceptance; Amount and Category are irrelevant.
- Each returned function retains its own supplied period; later constructor calls or caller changes do not alter it.

Public/reference assertions (actual/expected diagnostics):

- Case01: `f:=subject.WithinDays(subject.DaysPeriod{From:1,To:15});if f==nil{t.Fatal("got nil predicate")};for _,x:=range []struct{day int;want bool}{{0,false},{1,true},{13,true},{15,true},{16,false}}{if got:=f(subject.Record{Day:x.day,Amount:3,Category:"a"});got!=x.want{t.Errorf("day%d got %v want %v",x.day,got,x.want)}}`
- Case02: `f:=subject.WithinDays(subject.DaysPeriod{From:2,To:2});if f==nil||!f(subject.Record{Day:2}){t.Error("single-day period rejected")}`
- Case03: `f:=subject.WithinDays(subject.DaysPeriod{From:3,To:1});if f==nil{t.Fatal("nil predicate")};for _,day:=range []int{0,1,2,3,4}{if f(subject.Record{Day:day}){t.Errorf("reversed range accepted day%d",day)}}`
- Case04: `p:=subject.DaysPeriod{From:1,To:2};a:=subject.WithinDays(p);p.From=9;p.To=10;b:=subject.WithinDays(p);if a==nil||b==nil{t.Fatal("nil predicates")};if !a(subject.Record{Day:1})||a(subject.Record{Day:9})||!b(subject.Record{Day:9})||b(subject.Record{Day:1}){t.Error("predicates lost their independent periods")}`
- Case05: `f:=subject.WithinDays(subject.DaysPeriod{From:1,To:2});if f==nil||!f(subject.Record{Day:1,Amount:-99,Category:""}){t.Error("non-Day fields affected acceptance")}`

Realistic required defect models:

- exclude-from
- exclude-to
- or-bounds
- normalize-reversed
- global-period
- amount-required

Retained source behavior: expenses selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| inclusive bounds | 01–02 | exclude-from; exclude-to; or-bounds |
| reversed range rejected | 03 | normalize-reversed |
| independent captured values | 04 | global-period |
| other fields irrelevant | 05 | amount-required |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Public checks cover inclusive and singleton periods, reversed bounds, independent captured values and ignoring non-Day fields.

Reviewed the prompt, API, actual checks, all 6 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
